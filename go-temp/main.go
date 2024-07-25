package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx     context.Context
	redisDB *redis.Client
)

func IsFull() {
	val, err := redisDB.Get(ctx, "cwq").Result()
	if err != nil && err != redis.Nil {
		log.Printf("failed to get value from Redis: %v", err)
		return
	}
	fmt.Print(val)
	return
}

func init() {
	ctx = context.Background()
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:15001",
		Password: "",
		DB:       0, //redis单节点默认提供16个数据库，编号0-15
	})
}

func main() {
	IsFull()
}
