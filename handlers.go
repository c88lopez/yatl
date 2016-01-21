package main

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
)

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex func
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	var todos Todos

	c, _ := redis.Dial("tcp", ":6379")
	defer c.Close()

	values, _ := redis.Values(c.Do("LRANGE", "todos", 0, -1))

	redis.ScanSlice(values, &todos)

	fmt.Printf("%v\n", todos)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(todos))
}

// TodoShow func
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
