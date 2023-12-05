package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6370",
	})
	ctx := context.Background()
	result, err := rdb.Get(ctx, "breadcrumb/hondy").Result()
	if err != nil {
		panic(err)
	}
	data := []map[string]string{}

	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}
