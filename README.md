MUX
====

[![GoDoc](https://godoc.org/github.com/gorilla/mux?status.svg)](https://godoc.org/github.com/forgoer/mux)
[![LICENSE](https://img.shields.io/github/license/forgoer/thinkgo.svg)](https://github.com/forgoer/mux/blob/master/README.md)

## Installation

The only requirement is the [Go Programming Language](https://golang.org/dl/)

```
go get -u github.com/forgoer/mux
```

## Quick start

```go
package main

import (
    "fmt"
	"http"

	"github.com/forgoer/mux"
)

func main() {
	m := mux.New()
	m.HandleFunc("/", HomeHandler)
	m.HandleFunc("/products", ProductsHandler)
	m.HandleFunc("/articles", ArticlesHandler)
	http.ListenAndServe(":8080", &m)
}
```