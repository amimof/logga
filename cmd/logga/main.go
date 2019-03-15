package main

import (
	"fmt"
	"github.com/amimof/logga/pkg/api"
	"github.com/amimof/logga/pkg/server"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/pflag"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	VERSION   string
	COMMIT    string
	BRANCH    string
	GOVERSION string

	host string
	port int

	kubeconfigPath string
)

func init() {
	pflag.StringVar(&host, "host", "0.0.0.0", "The host address on which to listen for the --port port")
	pflag.StringVar(&kubeconfigPath, "kubeconfig", fmt.Sprintf("%s%s", os.Getenv("HOME"), "/.kube/config"), "Absolute path to a kubeconfig file")
	pflag.IntVar(&port, "port", 8080, "the port to listen on for insecure connections, defaults to 8080")
	klog.SetOutput(ioutil.Discard) // Tell klog, which is used by client-go to log into /dev/null instead of file
}

func main() {
	showver := pflag.Bool("version", false, "Print version")

	// parse the CLI flags
	pflag.Parse()

	// Show version if requested
	if *showver {
		fmt.Printf("Version: %s\nCommit: %s\nBranch: %s\nGoVersion: %s\n", VERSION, COMMIT, BRANCH, GOVERSION)
		return
	}

	// Check if kubeconfig exists. Set to "" if not so that we can use in-cluster config instead
	if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
		log.Printf("kubeconfig '%s' not found. Falling back to in-cluster config", kubeconfigPath)
		kubeconfigPath = ""
	}

	// Create config from environment
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	// Creates the clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	a, err := api.NewAPI(client)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	// Namespace routes
	r.Path("/api/namespaces").
		HandlerFunc(a.GetNamespaces)
	r.Path("/api/namespaces/{namespace}").
		HandlerFunc(a.GetNamespace)

	// Pod Routes
	r.Path("/api/namespaces/{namespace}/pods").
		HandlerFunc(a.GetPods)
	r.Path("/api/namespaces/{namespace}/pods/{pod}").
		HandlerFunc(a.GetPod)
	r.Path("/api/namespaces/{namespace}/pods/{pod}/log").
		Queries("watch", "true").
		HandlerFunc(a.StreamPodLog)
	r.Path("/api/namespaces/{namespace}/pods/{pod}/log").
		HandlerFunc(a.GetPodLog)

	// Serve the UI
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/logga/web/")))

	s := server.NewServer()
	s.Port = port
	s.Host = host
	s.ReadTimeout = 60 * time.Second
	s.WriteTimeout = 60 * time.Second
	s.Handler = corsOpts.Handler(r)

	// Listen and serve!
	err = s.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
