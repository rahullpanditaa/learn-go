package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world from %s\n", r.URL.Path[1:])
}

func main() {
	// bind the handler against a route
	http.HandleFunc("/", handler)

	// start a server -> open a TCP socket that can accept
	// HTTP requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}
