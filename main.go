package main

import (
	"myGin"
	"net/http"
)

/* Mostly based on net/http package */

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
	r.GET("/", func(c *myGin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello myGin</h1>")
	})

	r.GET("/hello", func(c *myGin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *myGin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *myGin.Context) {
		c.JSON(http.StatusOK, myGin.H{"filepath": c.Param("filepath")})
	})

	// now implement context for Query + PostForm
	// Context: create HTTP response effectively, clean up after HTTP request
	r.POST("/login", func(c *myGin.Context) {
		c.JSON(http.StatusOK, myGin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// Assume no context and we need to return a JSON, here is what we need to do:
	//obj = map[string]interface{}{
	//	"name":     "shuheng mo",
	//	"password": "123456",
	//}
	//w.Header().Set("Content-Type", "application/json") // head info
	//w.WriteHeader(http.StatusOK)                       // status code
	//encoder := json.NewEncoder(w)                      // create json encoder
	//if err := encoder.Encode(obj); err != nil {
	//	http.Error(w, err.Error(), 500)
	//}

	/* when matching the path we used map */
	/* However, it only matches static route but not like a type e.g. /username/:name */
	/* Thus, we introduce Trie tree here */

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
