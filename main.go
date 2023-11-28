package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6370",
	})
	ctx := context.Background()
	get := rdb.Get(ctx, "badtest")
	fmt.Println(get.Val(), get.Err())
	val, err := rdb.Get(ctx, "test").Result()
	fmt.Println(val, err)
}
