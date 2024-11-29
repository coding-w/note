package main

// 基础数据类型常规用法
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var (
	client redis.UniversalClient
	ctx    = context.Background()
	wg     sync.WaitGroup
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 1. String Demo
	//stringDemo()

	// 2. List Demo
	//listDemo()

	// 3. Hash Demo
	//hashDemo()

	// 4. Set Demo
	//setDemo()

	// 5. ZSet Demo
	//zsetDemo()

	// 6. bitmap
	//bitmapDemo()

	// 7. Geo
	//geoDemo()

	// 8. Stream()
	streamDemo()

	//clearAll()
}

func streamDemo() {
	producer := func() {
		defer wg.Done()
		// 向 Redis Stream 中添加任务
		// 使用 "*" 自动生成消息 ID
		taskID := "task123"
		data := map[string]interface{}{
			"task_id":     taskID,
			"description": "Process some data",
			"priority":    "high",
		}

		// 将任务添加到 Stream（名为 "taskQueue"）
		_, err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "taskQueue",
			Values: data,
		}).Result()
		if err != nil {
			log.Fatalf("Error adding task to stream: %v", err)
		}
		fmt.Printf("Task %s added to the stream\n", taskID)
	}

	consumer := func(groupName, consumerName string) {
		defer wg.Done()
		// 创建消费组（如果还不存在）
		// XGroupCreateMkStream：如果消费组已经存在，不会抛出错误
		//_, err := client.XGroupCreateMkStream(ctx, "taskQueue", groupName, "$").Result()
		//if err != nil && err != redis.Nil {
		//	log.Fatalf("Error creating consumer group: %v", err)
		//} else {
		//	fmt.Printf("Consumer group %s created or already exists.\n", groupName)
		//}

		// 消费者从流中读取任务
		for {
			// 使用 XREADGROUP 来读取任务
			// "$" 表示从未确认的消息开始读取
			msgs, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    groupName,
				Consumer: consumerName,
				Streams:  []string{"taskQueue", ">"},
				Block:    0, // 阻塞等待直到有消息可读
				Count:    1, // 每次读取一条消息
			}).Result()
			if err != nil {
				log.Printf("Error reading from stream: %v", err)
				continue
			}

			// 处理任务
			for _, msg := range msgs {
				for _, xMessage := range msg.Messages {
					taskID := xMessage.Values["task_id"].(string)
					description := xMessage.Values["description"].(string)
					priority := xMessage.Values["priority"].(string)

					// 任务处理逻辑（示例：简单打印任务信息）
					fmt.Printf("[%s] Consumer %s processing task %s: %s (Priority: %s)\n", consumerName, groupName, taskID, description, priority)

					// 确认消息已处理
					_, err := client.XAck(ctx, "taskQueue", groupName, xMessage.ID).Result()
					if err != nil {
						log.Printf("Error acknowledging task %s: %v", taskID, err)
					} else {
						fmt.Printf("Task %s acknowledged\n", taskID)
					}
				}
			}
		}
	}
	wg.Add(2)
	// 启动生产者将任务添加到 Redis Stream
	go producer()

	// 启动两个消费者处理任务
	go consumer("taskGroup", "consumer1")
	//go consumer("taskGroup", "consumer2")

	// 等待消费者处理任务
	wg.Wait()
}

func geoDemo() {
	// Geo 数据类型
	client.GeoAdd(ctx, "geo:location:user", &redis.GeoLocation{Name: "zs", Longitude: 116.044567, Latitude: 39.030452})
	client.GeoAdd(ctx, "geo:location:car", &redis.GeoLocation{Name: "1", Longitude: 116.034767, Latitude: 39.030352})
	r1 := client.GeoPos(ctx, "geo:location:user", "ww").Val()
	r2, _ := client.GeoPos(ctx, "geo:location:user", "zs").Result()
	fmt.Println(r1[0].Latitude, r1[0].Longitude)
	fmt.Println(r2[0].Latitude, r2[0].Longitude)
	val := client.GeoDist(ctx, "geo:location:user", "ww", "zs", "km").Val()
	fmt.Println(val)
	locations := client.GeoRadius(ctx, "geo:location:car", 116.044567, 39.030452, &redis.GeoRadiusQuery{Radius: 1, Unit: "km"}).Val()
	fmt.Println(locations)
}

func bitmapDemo() {
	// 打卡签到统计
	client.SetBit(ctx, "user:1:202412", 1, 1)
	client.SetBit(ctx, "user:1:202412", 2, 0)
	client.SetBit(ctx, "user:1:202412", 3, 1)
	client.SetBit(ctx, "user:2:202412", 1, 0)
	client.SetBit(ctx, "user:2:202412", 2, 1)
	client.SetBit(ctx, "user:2:202412", 3, 1)
	// 有多少个 1
	val := client.BitCount(ctx, "user:1:202412", &redis.BitCount{Start: 0, End: -1}).Val()
	fmt.Println(val)
	result, _ := client.BitOpNot(ctx, "user:1:202412", "user:2:202412").Result()
	fmt.Println(result)
}

func zsetDemo() {
	// 有序集合
	result, _ := client.ZAdd(ctx, "article:list", redis.Z{Score: 55, Member: "article:1"}, redis.Z{Score: 54, Member: "article:2"}, redis.Z{Score: 56, Member: "article:3"}, redis.Z{Score: 57, Member: "article:4"}).Result()
	fmt.Println(result)
	list, _ := client.ZRangeWithScores(ctx, "article:list", 0, -1).Result()
	fmt.Println(list)
	// 增加文章的阅读量
	f, _ := client.ZIncrBy(ctx, "article:list", 100, "article:1").Result()
	fmt.Println(f)
	res, _ := client.ZRevRangeWithScores(ctx, "article:list", 0, -1).Result()
	fmt.Println(res)

}

func setDemo() {
	base := func() {
		// 常见的操作
		// 添加元素
		_, err := client.SAdd(ctx, "list", "zs", "li", "ww", "sdfghjkl;").Result()
		if err != nil {
			log.Fatal(err)
		}
		// 获取元素
		result, err := client.SMembers(ctx, "list").Result()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		// 列表数量
		val := client.SCard(ctx, "list").Val()
		fmt.Println(val)
		// 删除指定元素
		rem := client.SRem(ctx, "list", "zs")
		fmt.Println(rem)
		// 随机删除
		strings, _ := client.SPopN(ctx, "list", 1).Result()
		fmt.Println(strings)
	}
	base()

	// 用户点赞
	case1 := func() {
		// key
		key := "article:1"
		client.SAdd(ctx, key, "uid:1")
		client.SAdd(ctx, key, "uid:2")
		client.SAdd(ctx, key, "uid:3")
		client.SAdd(ctx, key, "uid:4")
		// uid:1 有没有点赞
		member := client.SIsMember(ctx, key, "uid:1")
		fmt.Println("uid:1 点赞了吗", member)
		// 获取点赞数量
		i := client.SCard(ctx, key).Val()
		fmt.Println("article:1 点赞数量", i)
		// 用户uid:3 取消点赞
		client.SRem(ctx, key, "uid:3")
		// 点赞的列表
		res, _ := client.SMembers(ctx, key).Result()
		fmt.Println(res)
	}
	case1()

	// 集合交、并、差运算
	setOper := func() {
		key1 := "list1"
		key2 := "list2"
		client.SAdd(ctx, key1, 1, 2, 3, 4)
		client.SAdd(ctx, key2, 3, 4, 5, 6)
		result, _ := client.SInterStore(ctx, "list3", key1, key2).Result()
		fmt.Println("集合并集并保存", result)
		res, _ := client.SMembers(ctx, key1).Result()
		fmt.Println("集合并集结果", res)
		union := client.SUnion(ctx, key1, key2).Val()
		fmt.Println("集合交集结果", union)
		diff := client.SDiff(ctx, key1, key2).Val()
		fmt.Println("集合差集结果", diff)
	}
	setOper()
}

func hashDemo() {
	// 存储对象
	type stu struct {
		Name string
		Age  int
	}

	_, err := client.HSet(ctx, "stu", "name", "zs", "age", 22, "xl", "whxy").Result()
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.HGetAll(ctx, "stu").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

func listDemo() {
	// 简单插入和读取
	key1 := "list"
	err := client.LPush(ctx, key1, "wxasd").Err()
	if err != nil {
		log.Fatal(err)
	}
	results, err := client.LRange(ctx, key1, 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
	// 消息队列 右侧插入，左侧读取
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err = client.RPush(ctx, "queue", fmt.Sprintf("%d", i)).Err()
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			val := client.BLPop(ctx, 0, "queue").Val()
			fmt.Println(val)
		}
	}()
	wg.Wait()

}

func stringDemo() {
	count := client.Exists(ctx, "foo").Val()
	if count > 0 {
		fmt.Println("foo is exists")
		return
	}
	// 基础的缓存
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("foo", val)
	// 实现文章阅读计数器
	key := "aritcle:readcount:1001"
	client.Set(ctx, key, 0, 0)
	for i := 0; i < 10; i++ {
		client.Incr(ctx, key)
	}
	val, err = client.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("aritcle:readcount:1001", val)
}

func clearAll() {
	result, err := client.FlushAll(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
