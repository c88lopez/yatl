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
	// var todos string
	// todos := Todos{
	// 	Todo{Name: "Write presentation"},
	// 	Todo{Name: "Host meetup"},
	// }
	//
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }

	c, _ := redis.Dial("tcp", ":6379")
	defer c.Close()

	c.Do("SET", "todos", `{task: "task1"}`)

	todos, _ := redis.String(c.Do("GET", "todos"))
	// redis.Scan(reply, &todos)

	fmt.Fprint(w, todos)

	// fmt.Printf("%#v\n", todos)

	// dec := json.NewDecoder(strings.NewReader(todos))
	// fmt.Printf("%v\n", reply)

	// fmt.Printf("%#v\n", dec)
}

// TodoShow func
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
