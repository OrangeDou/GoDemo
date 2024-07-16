package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	ctx     context.Context
	redisDB *redis.Client
)

func ConnectRedis() {

	val, err := redisDB.Do(ctx, "incr", "api").Result()
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(val)

}

func IsFull() (bool, error) {
	val, err := redisDB.Get(ctx, "api").Result()
	if err != nil && err != redis.Nil {
		log.Printf("failed to get value from Redis: %v", err)
		return false, err
	}

	// 如果键不存在，返回 0
	valInt := 0
	if val != "" {
		valInt, err = strconv.Atoi(val)
		if err != nil {
			log.Printf("failed to convert '%s' to integer: %v", val, err)
			return false, err
		}
	}

	return valInt < 200, nil
}

func init() {
	ctx = context.Background()
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, //redis单节点默认提供16个数据库，编号0-15
	})
}
