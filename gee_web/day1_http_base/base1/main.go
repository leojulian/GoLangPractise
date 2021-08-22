package main

import (
	"fmt"
	"log"
	"net/http"
)

// curl http://localhost:8080
// curl http://localhost:8080/Hello
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path is %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
