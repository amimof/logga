package api

import (
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
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
}

func NewAPI() (*API, error) {

	c, err := makeClient()
	if err != nil {
		return nil, err
	}
	return &API{
		client: c,
	}, nil
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

// GetPods will return PodList in a namespace
func (a *API) GetPods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	statusCode := 200

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
		http.Error(w, res.Error().Error(), statusCode)
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
	vars := mux.Vars(r)
	statusCode := 200

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
		http.Error(w, res.Error().Error(), statusCode)
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
	statusCode := 200

	res := a.client.
		CoreV1().
		RESTClient().
		Get().
		Resource("namespaces").
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
		http.Error(w, res.Error().Error(), statusCode)
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
	vars := mux.Vars(r)
	statusCode := 200

	res := a.client.
		CoreV1().
		RESTClient().
		Get().
		Resource("namespaces").
		Name(vars["namespace"]).
		Do().
		StatusCode(&statusCode)

	if res.Error() != nil {
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
		http.Error(w, res.Error().Error(), statusCode)
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

	buf := make([]byte, 1024*16)
	canaryLog := []byte("unexpected stream type \"\"")

	for ctx.Err() == nil {
		nread, err := stream.Read(buf)

		switch {
		case err == io.EOF:
			log.Printf("%s", err.Error())
			return
		case ctx.Err() != nil:
			log.Printf("%s", ctx.Err())
			return
		case err != nil:
			log.Printf("%s", err.Error())
			return
		case nread == 0:
			log.Printf("%s", io.EOF)
			return
		}

		l := buf[0:nread]

		if bytes.Compare(canaryLog, l) == 0 {
			log.Printf("received 'unexpect stream type'")
			continue
		}

		_, err = w.Write(l)
		if err != nil {
			log.Printf("%s", err.Error())
			break
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

	}

}
