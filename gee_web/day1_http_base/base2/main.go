package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path is %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not found: %q\n", req.URL.Path)
	}
}

// curl http://localhost:8080
// curl http://localhost:8080/hello
// curl http://localhost:8080/hello123
func main() {
	engine := new(Engine)
	http.ListenAndServe(":8080", engine)
}
