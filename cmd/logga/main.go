package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/pflag"
	"gitlab.com/amimof/logga/pkg/api"
	"gitlab.com/amimof/logga/pkg/server"
	"log"
	"net/http"
)

var (
	VERSION   string
	COMMIT    string
	BRANCH    string
	GOVERSION string
)

func main() {

	showver := pflag.Bool("version", false, "Print version")

	// parse the CLI flags
	pflag.Parse()

	// Show version if requested
	if *showver {
		fmt.Printf("Version: %s\nCommit: %s\nBranch: %s\nGoVersion: %s\n", VERSION, COMMIT, BRANCH, GOVERSION)
		return
	}

	a, err := api.NewAPI()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	corsOpts := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:8081"}, //you service is available and allowed for this base url 
    AllowedMethods: []string{
        http.MethodGet,//http methods for your app
        http.MethodPost,
        http.MethodPut,
        http.MethodPatch,
        http.MethodDelete,
        http.MethodOptions,
        http.MethodHead,
		},
    AllowedHeaders: []string{
			"*",//or you can your header key values which you are using in your application
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

	s := server.NewServer()
	s.Handler = corsOpts.Handler(r)

	// Listen and serve!
	err = s.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
