package main

import (
	"fmt"
	"myGin"
	"net/http"
)

/* Mostly based on net/http package */

/*
The project structure
myGin/
|--gin.go
|--go.mod
main.go
go.mod
*/

/*
The source code:
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error

h Handler --> an interface
to make it handle all the HTTP request, we will need to pass
in an instance that implemented ServeHTTP()
*/

//type Engine struct {
//}
//
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	case "/hello":
//		for k, v := range req.Header {
//			fmt.Fprintf(w, "Header[%q] =  %q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
//	}
//}

func main() {
	// Dummy example
	// https.HandlerFunc(): router path and self-defined functions
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/hello", helloHandler)
	// start the service
	// engine := new(Engine)
	// log.Fatal(http.ListenAndServe(":9999", engine))

	r := myGin.New()
	// register the methods and then run
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}

/*
2 arguments to maintain in handler functions:
w http.ResponseWriter
req *http.Request
*/

//func indexHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//}
//
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "Header[%q] =  %q\n", k, v)
//	}
//}
