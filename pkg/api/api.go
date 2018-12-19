package api

import (
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"context"
	"log"
	"time"
	"bytes"
)

type API struct {

}

func NewAPI() *API {
	return &API{}
}

func (a *API) makeClient(ctx context.Context) (corev1.CoreV1Interface, error) {
	c, err := clientcmd.BuildConfigFromFlags("", "/Users/amir/.kube/config")
	if err != nil {
		return nil, err
	}
	cs, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return cs.CoreV1(), nil
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	log.Printf("Getting logs ...")

	ctx, cancel := context.WithCancel(r.Context())
	client, err := a.makeClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cancel()
		return
	}

	since := int64(time.Second * 60)
	opts := &v1.PodLogOptions{
		Container:    "busybox",
		Follow:       true,
		SinceSeconds:	&since,
	}	

	req := client.
		Pods("default").
		GetLogs("busybox-57b4f54479-c5w74", opts).
		Context(ctx)

	stream, err := req.Stream()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer stream.Close()
	
	buf := make([]byte, 1024 * 16)
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

	log.Printf("Done")

}