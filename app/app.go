package app

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisServer *redis.Client

// InitRedis 获得Redis实例
func InitRedis() (*redis.Client, error) {
	redisServer := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 检测心跳
	pong, err := redisServer.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect redis failed")
		return nil, err
	}
	fmt.Printf("redis ping result: %s\n", pong)
	return redisServer, nil
}

// Destruct 销毁
func Destruct() {
	if redisServer != nil {
		_ = redisServer.Close()
	}
}
