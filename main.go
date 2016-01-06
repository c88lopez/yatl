package main

import (
	"log"
	"net/http"
)

// main func
func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
