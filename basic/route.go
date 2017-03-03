package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "{\"content\":\"Index\"}")
}

func show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "{\"content\":\"Show Detail\"}")
}

func registerRoute() {
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	registerRoute()
}
