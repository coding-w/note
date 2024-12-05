package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

// 集群的使用
var (
	redisClient *redis.ClusterClient
	ctx         = context.Background()
	wg          sync.WaitGroup
)

func init() {
	redisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"192.168.3.222:7000", "192.168.3.222:7001", "192.168.3.222:7002"},
		Password: "",
	})
	// 测试连接
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis cluster: %v", err)
	}
	fmt.Println("Successfully connected to Redis cluster!")
}

func main() {
	key := "example_key"
	value := "hello_redis_cluster"

	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set key: %v", err)
	}

	result, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}
	fmt.Printf("Got value for '%s': %s\n", key, result)

	// 关闭连接
	err = redisClient.Close()
	if err != nil {
		log.Fatalf("Failed to close Redis cluster client: %v", err)
	}
}
