package api

import (
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"fmt"
	//corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/rest"
	"bytes"
	"context"
	"github.com/gorilla/mux"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"strconv"
	"time"
)

type API struct {
	client *kubernetes.Clientset
	broker *Broker
}

type Broker struct {
	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

func NewAPI() (*API, error) {

	c, err := makeClient()
	if err != nil {
		return nil, err
	}
	return &API{
		client: c,	
		broker: NewBroker(),
	}, nil
	
}

func NewBroker() *Broker {
	broker := &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}
	go broker.listen()
	return broker
}

func makeClient() (*kubernetes.Clientset, error) {
	c, err := clientcmd.BuildConfigFromFlags("", "/Users/amir/.kube/config")
	if err != nil {
		return nil, err
	}
	cs, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (broker *Broker) listen() {
	patience := time.Second*1
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan, _ := range broker.clients {
				select {
				case clientMessageChan <- event:
				case <-time.After(patience):
					//log.Print("Skipping client.")
					continue
				}
			}
		}
	}

}

// GetPods will return PodList in a namespace
func (a *API) GetPods(w http.ResponseWriter, r *http.Request) {
	
	var statusCode int
	vars := mux.Vars(r)

	res := a.client.
		CoreV1().
		RESTClient().
		Get().
		SetHeader("Accept", "application/json").
		Namespace(vars["namespace"]).
		Resource("pods").
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, res.Error().Error(), statusCode)
		return
	}

	b, err := res.Raw()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

// GetPods will return a Pod in a namespace
func (a *API) GetPod(w http.ResponseWriter, r *http.Request) {
	
	var statusCode int
	vars := mux.Vars(r)

	res := a.client.
		CoreV1().
		RESTClient().
		Get().
		SetHeader("Accept", "application/json").
		Namespace(vars["namespace"]).
		Resource("pods").
		Name(vars["pod"]).
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, res.Error().Error(), statusCode)
		return
	}

	b, err := res.Raw()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

// Namespaces returns NamespaceList in a cluster
func (a *API) GetNamespaces(w http.ResponseWriter, r *http.Request) {

	var statusCode int

	res := a.client.
		CoreV1().
		RESTClient().
		Get().
		Resource("namespaces").
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, res.Error().Error(), statusCode)
		return
	}

	b, err := res.Raw()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

// GetNamespace returns a Namespace in a cluster
func (a *API) GetNamespace(w http.ResponseWriter, r *http.Request) {
	
	var statusCode int
	vars := mux.Vars(r)

	res := a.client.
		CoreV1().
		RESTClient().
		Get().
		Resource("namespaces").
		Name(vars["namespace"]).
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, res.Error().Error(), statusCode)
		return
	}

	b, err := res.Raw()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

func sinceSeconds(s string) int64 {
	since := int64(1)
	if s == "" {
		return since
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return since
	}
	return int64(time.Second * time.Duration(i))
}

func tailLines(s string) int64 {
	tail := int64(1000)
	if s == "" {
		return tail
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return tail
	}
	return int64(i)
}

// GetPodLog returns container log of a container in a pod
func (a *API) GetPodLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	container := r.URL.Query().Get("container")
	statusCode := 200
	//since := sinceSeconds(r.URL.Query().Get("sinceSeconds"))
	tail := tailLines(r.URL.Query().Get("tailLines"))

	opts := &v1.PodLogOptions{
		Container: container,
		//SinceSeconds: &since,
		TailLines: &tail,
	}

	res := a.client.
		CoreV1().
		Pods(vars["namespace"]).
		GetLogs(vars["pod"], opts).
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, res.Error().Error(), statusCode)
		return
	}

	b, err := res.Raw()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)

}

// StreamPodLog will start streaming logs of a container running in a pod
func (a *API) StreamPodLog(w http.ResponseWriter, r *http.Request) {

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not unsupported", http.StatusInternalServerError)
		return
	}

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	a.broker.newClients <- messageChan
	
	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		a.broker.closingClients <- messageChan
	}()

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	container := r.URL.Query().Get("container")
	ctx, cancel := context.WithCancel(r.Context())
	since := sinceSeconds(r.URL.Query().Get("sinceSeconds"))

	opts := &v1.PodLogOptions{
		Container:    container,
		Follow:       true,
		SinceSeconds: &since,
	}

	req := a.client.
		CoreV1().
		Pods(vars["namespace"]).
		GetLogs(vars["pod"], opts).
		Context(ctx)

	stream, err := req.Stream()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cancel()
		return
	}

	defer stream.Close()

	// Listen to connection close and un-register messageChan
	notify := w.(http.CloseNotifier).CloseNotify()

	buf := make([]byte, 1024*16)
	canaryLog := []byte("unexpected stream type \"\"")

	for {
		if(ctx.Err() != nil) {
			//log.Printf("Erro: %s", ctx.Err())
			return
		}
		select {
		case <-ctx.Done():
			return
		case <-notify:
			return
		default:
			nread, err := stream.Read(buf)

			switch {
			case err == io.EOF:
				//log.Printf("%s", err.Error())
				return
			case err != nil:
				//log.Printf("nil error %s", err.Error())
				return
			case nread == 0:
				//log.Printf("%s", io.EOF)
				return
			default:
				l := buf[0:nread]
	
				if bytes.Compare(canaryLog, l) == 0 {
					log.Printf("received 'unexpect stream type'")
					continue
				}
				a.broker.Notifier <- l
				fmt.Fprintf(w, "data: %s\n\n", <-messageChan)
				flusher.Flush()
			}
	
		}

	}

}
