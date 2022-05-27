package main

import (
	"fmt"
	"log"
	"net/http"
)

var hits int =0

func handler(w http.ResponseWriter, r *http.Request) {
	hits++
	fmt.Fprintf(w, "Hello World!")
}

func queries(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Total number of hits: %d",hits)
}

func main() {
	http.HandleFunc("/requests", handler)
	http.HandleFunc("/queries", queries)
	log.Fatal(http.ListenAndServe(":8068",nil))
}