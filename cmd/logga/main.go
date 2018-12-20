package main 

import (
	"log"
	"gitlab.com/amimof/logga/pkg/server"
	"gitlab.com/amimof/logga/pkg/api"
	"github.com/gorilla/mux"
	//"net/http"
)

func main() {

	a, err := api.NewAPI()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	//r.HandleFunc(Path("/").Handler(http.FileServer(http.Dir("pkg/http/web")))
	r.HandleFunc("/api/namespaces", a.GetNamespaces)
	r.HandleFunc("/api/namespaces/{namespace}", a.GetNamespace)
	r.HandleFunc("/api/namespaces/{namespace}/pods", a.GetPods)
	r.HandleFunc("/api/namespaces/{namespace}/pods/{pod}", a.GetPod)
	r.HandleFunc("/api/namespaces/{namespace}/pods/{pod}/stream", a.StreamPodLog)
	
	s := server.NewServer()
	s.Handler = r

	// Listen and serve!
	err = s.Serve()
	if err != nil {
		log.Fatal(err)
	}

}