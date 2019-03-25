# logga
[![Build Status](https://travis-ci.com/amimof/logga.svg?token=YU8cQELmfms9zTY3ztML&branch=master)](https://travis-ci.com/amimof/logga) [![logga](https://godoc.org/github.com/amimof/logga?status.svg)](https://godoc.org/github.com/amimof/logga) [![Go Report Card](https://goreportcard.com/badge/github.com/amimof/logga)](https://goreportcard.com/report/github.com/amimof/logga) [![Coverage](http://gocover.io/_badge/github.com/amimof/logga)](http://gocover.io/github.com/amimof/logga)

----

With `logga` you view [Kubernetes](https://kubernetes.io) container logs using a web browser. Sounds simple right? That's the point. It's easy to use and designed to run on [Kubernetes](https://kubernetes.io). It's used by application teams, developers and sysadmins to effortlessly view container logs using a web browser, without a complex logging infrastructure. 

## Key features

* Super easy do deploy and manage
* Fully stateless 
* In-cluster as well as out-of-cluster support
* Never leave the keyboard with quick commands
* Quick search
* Tail/Watch logs

![](./img/logga_screenshot.png)

## Getting started

### In-cluster (recommended)
This is the prefered method of deploying and running. Logga is built to run on Kubernetes.

```
kubectl apply -f https://raw.githubusercontent.com/amimof/logga/master/deploy/kubernetes.yml
```

### Out-of-cluster
It is possible to run logga out-of-cluster using docker. Logga is able to use credentials to establish a connection to a cluster using a kubernetes [kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/).
```
docker run \
  -d \
  --name logga \
  -v $PWD/.kube/config:/config \
  -p 8080:8080 amimof/logga:latest \
  --kubeconfig=/config
```

## Contributing

All help in any form is highly appreciated and your are welcome participate in developing `logga` together. To contribute submit a `Pull Request`. If you want to provide feedback, open up a Github `Issue` or contact me personally. 