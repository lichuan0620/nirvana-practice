# Nirvana Practice Project

As a HTTP server framework, [Nirvana](https://github.com/caicloud/nirvana) is robust but can also be quite complicated. This project let you practice writing Restful APIs with Nirvana.

## Prerequisites

 - [Golang](https://golang.org/dl/) - the Go language
 - [dep](https://github.com/golang/dep) - the Go dependency management tool used by Nirvana
 
That is. You'll use dep to install the required Go packages.

## Getting Started

The project is already runnable as it is. Use dep to install the dependencies and the Makefile to build it.

```
$ dep ensure -v

$ VERSION=v0.1.0 make build
```

Check out the help message and version information.

```
 $ bin/practice-server -h
Usage of bin/practice-server:
  -p, --port uint16   the HTTP port used by the server (default 8080)
  -v, --version       show version info
pflag: help requested

$ bin/nirvana-practice -v
nirvana-practice, version v0.1.0 (branch: master), revision: ab9132625bd5e2b9c8d3b2ff23c423cc9c02e9ec)
```

Now run the server and try the APIs.

```
$ bin/practice-server -p 8080

$ curl localhost:8080/api/v1alpha1/products/null
{"reason":"practice:NotImplemented","message":"requested feature is not implemented"}

$ curl localhost:8080/api/v1alpha1/products
{"reason":"Nirvana:Router:RouterNotFound","message":"can't find router"}
```

## Do It Yourself

The product API is already defined in `pkg/apis/v1alpha1/products.go` for you, alone with an incomplete `GET` API. Try to finish the `GET` API and write `POST`, `LIST`, `PUT`, and `DELETE` as well. 

Use GitHub issue to ask questions and, once finished, submit a PR to this repo to show off your result. 

## Containerize Your App

Once your code is finished and tested, you might want to containerize your app so that you can deploy it practically. The Dockerfile is created for you under the `/build` directory. Try to finish it run `make container` to build the container images. 

Check out [the official Docker Getting-Started guide](https://docs.docker.com/get-started/) if you are not already familiar with Docker and containers. 
