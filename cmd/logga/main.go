package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	"gitlab.com/amimof/logga/pkg/api"
	"gitlab.com/amimof/logga/pkg/server"
	"log"
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
	//r.HandleFunc(Path("/").Handler(http.FileServer(http.Dir("pkg/http/web")))

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
	s.Handler = r

	// Listen and serve!
	err = s.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
