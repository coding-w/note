package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

// 对Redis的主从的使用
var (
	redisClient *redis.Client
	ctx         = context.Background()
	wg          sync.WaitGroup
)

func init() {
	// 配置哨兵x
	redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",                                                        // 主节点的名字，必须和哨兵配置一致
		SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"}, // 哨兵地址列表
		Password:      "123456",                                                          // 如果有密码需要配置
		DB:            0,                                                                 // 使用的数据库编号
		PoolSize:      10,                                                                // 连接池大小
	})
	// 测试连接
	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis via Sentinel: %v", err))
	}
	fmt.Println("Connected to Redis via Sentinel.")
}

// 写数据
func writeToRedis(ctx context.Context, key, value string, expiration time.Duration) error {
	err := redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 读数据
func readFromRedis(ctx context.Context, key string) (string, error) {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func main() {

	defer redisClient.Close()
	// 写入数据
	err := writeToRedis(ctx, "example-key", "example-value", 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to write to Redis: %v\n", err)
		return
	}
	fmt.Println("Data written successfully.")

	time.Sleep(25 * time.Second)

	// 读取数据
	value, err := readFromRedis(ctx, "example-key")
	if err != nil {
		fmt.Printf("Failed to read from Redis: %v\n", err)
		return
	}
	fmt.Printf("Data read successfully: %s\n", value)
}
