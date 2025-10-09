package main

// solving data race using channels
import (
	"fmt"
	"log"
	"net/http"
)

var nextId int

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d<h1>", nextId)

	nextId++ // UNSAFE
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
