package main 

import (
	"log"
	"gitlab.com/amimof/logga/pkg/server"
	"gitlab.com/amimof/logga/pkg/api"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	a := api.NewAPI()

	r := mux.NewRouter()
	r.Path("/").Handler(http.FileServer(http.Dir("pkg/http/web")))
	r.Path("/api").Handler(a)
	
	s := server.NewServer()
	s.Handler = r

	// Listen and serve!
	err := s.Serve()
	if err != nil {
		log.Fatal(err)
	}

}