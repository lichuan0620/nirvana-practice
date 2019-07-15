# Nirvana Practice Project

As a HTTP server framework, [Nirvana](https://github.com/caicloud/nirvana) is robust but can also be quite complicated. This project let you learn writing Restful APIs with Nirvana with practice.

## Prerequisites

 - [Golang](https://golang.org/dl/) - the Go language  
 - [Nirvana](https://github.com/caicloud/nirvana) - hopefully you know what your are learning
 - [dep](https://github.com/golang/dep) - the Go dependency management tool used by Nirvana

## Getting Started

The project is already runnable as it is. Use dep to install the dependencies and the Makefile to build it.

```
$ dep ensure -v
$ VERSION=v0.1.0 make build
```

Check out the help message and version information.

```
$ bin/practice-server -h

ENV-Flag Mapping Table

   ENV        Flag         Current Value  
1  HTTP_PORT  --http-port  8080           
2  VERSION    --version    false

Usage:
  alerting-admin [flags]

Flags:
  -h, --help               help for alerting-admin
  -p, --http-port uint16   port on which the HTTP server would be exposed (default 8080)
  -v, --version            show version

$ bin/nirvana-practice -v
nirvana-practice, version v0.1.0 (branch: master), revision: HEAD)
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

The product API is already defined in `pkg/apis/v1alpha1/products.go` for you, alone with an incomplete `GET` API. Try to finish the `GET` API and write `POST`, `LIST`, `PUT`, and `DELETE` as well. Use GitHub issue to ask questions and, once finished, submit a PR to this repo to showoff your result. 
