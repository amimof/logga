package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/http"
	"strconv"
	"time"
)

// API exposes client and broker and various methods to interact with backend api-servers
type API struct {
	client *kubernetes.Clientset
	broker *Broker
}

// NewAPI returns an initialized API instance
func NewAPI(c *kubernetes.Clientset) (*API, error) {
	return &API{
		client: c,
		broker: NewBroker(),
	}, nil

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

// GetPod will return a Pod in a namespace
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

// GetNamespaces returns NamespaceList in a cluster
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

	defer cancel()

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
		return
	}

	defer stream.Close()

	// Listen to connection close and un-register messageChan
	notify := w.(http.CloseNotifier).CloseNotify()

	buf := make([]byte, 1024*16)
	canaryLog := []byte("unexpected stream type \"\"")

	for {
		if ctx.Err() != nil {
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
				return
			case err != nil:
				return
			case nread == 0:
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
