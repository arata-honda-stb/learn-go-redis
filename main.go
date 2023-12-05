package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func getBreadCrumbs(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6370",
	})
	ctx := context.Background()
	key := r.URL.Path[len("/breadcrumb/"):]
	result, err := rdb.Get(ctx, "breadcrumb/"+key).Result()
	if err != nil {
		panic(err)
	}
	data := []map[string]string{}
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/breadcrumb/", getBreadCrumbs)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
