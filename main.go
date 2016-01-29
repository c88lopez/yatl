package main

import (
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

// main func
func main() {
	redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			return nil, err
		}

		return c, err
	}, 1000)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
