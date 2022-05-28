package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

func handler(w http.ResponseWriter, r *http.Request) {
	// incrementRedis(hits)
	client.Incr(ctx, "hits")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Fprintf(w, "Hello World!")
}

func queries(w http.ResponseWriter, r *http.Request) {
	val, err := client.Get(ctx, "hits").Result()
	fmt.Fprintf(w, "Total number of hits: %s",val)
	fmt.Print(err)
}

// func incrementRedis(hits int) {
// 	err := client.Set(ctx, "hits", hits, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	http.HandleFunc("/requests", handler)
	http.HandleFunc("/queries", queries)
	log.Fatal(http.ListenAndServe(":8068",nil))
}