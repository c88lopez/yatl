package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Todo struct
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos list
type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
