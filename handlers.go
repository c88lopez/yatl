package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Fprintln(w, "Welcome!")

	fmt.Printf("%v\n", time.Since(now))
}

// TodoIndex func
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	//c, _ := redis.Dial("tcp", ":6379")
	c := redisPool.Get()
	defer c.Close()

	todos, _ := redis.String(c.Do("GET", "todos"))

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(todos))

	fmt.Printf("%v\n", time.Since(now))
}

// TodoSave func
func TodoSave(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("%#v", r.PostFormValue("todos"))

	c, _ := redis.Dial("tcp", ":6379")
	defer c.Close()

	_, err := c.Do("SET", "todos", r.PostFormValue("todos"))

	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}

	fmt.Printf("%v\n", time.Since(now))
}
