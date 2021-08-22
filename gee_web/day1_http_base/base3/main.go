package main

import (
	"fmt"
	"gee"
	"net/http"
)

// $curl http://localhost:8080
// $curl http://localhost:8080/hello
// $curl http://localhost:8080/hello12
func main() {

	g := gee.New()

	g.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Pat is %q\n", req.URL.Path)
	})

	g.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", v, k)
		}
	})

	g.Run(":8080")
}
