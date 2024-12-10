package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"note/grpc/stream_bidirectional/example"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := example.NewExampleServiceClient(conn)
	var globalWg sync.WaitGroup
	globalWg.Add(1)
	go func() {
		defer globalWg.Done()
		stream, err := client.Example(context.Background())
		if err != nil {
			log.Fatalf("Failed to call Example: %v", err)
		}
		var wg sync.WaitGroup
		wg.Add(2)
		// 启动发送和接收 goroutine
		go sendMessages(stream, &wg)
		go receiveMessages(stream, &wg)
		// 等待发送和接收的 goroutine 完成
		wg.Wait()
	}()

	// Handle ExampleForever stream (server streaming)
	//globalWg.Add(1)
	//go handleExampleForever(client, &globalWg)

	// 等待所有 goroutine 完成
	globalWg.Wait()
}

// sendMessages 处理客户端流的发送逻辑
func sendMessages(stream example.ExampleService_ExampleClient, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		err := stream.Send(&example.ExampleRequest{
			Title: fmt.Sprintf("[Example]Client Stream Send msg index:[%d]\n", i),
		})
		if err != nil {
			if err == io.EOF {
				fmt.Println("bye...")
				break
			}
			log.Printf("Failed to send message: %v", err)
			return
		}

		if i == 5 {
			fmt.Println("apply close.")
			if err := stream.CloseSend(); err != nil {
				log.Printf("Failed to close send stream: %v", err)
			}
			break
		}

		time.Sleep(time.Second)
	}
}

// receiveMessages 处理服务器流的接收逻辑
func receiveMessages(stream example.ExampleService_ExampleClient, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("bye...")
				break
			}
			log.Printf("Failed to receive message: %v", err)
			return
		}
		fmt.Printf("[Example]receive server msg: %s\n", msg.Reply)
	}
}

// handleExampleForever 处理 ExampleForever 流式 RPC
func handleExampleForever(client example.ExampleServiceClient, wg *sync.WaitGroup) {
	defer wg.Done()
	stream, err := client.ExampleForever(context.Background(), &example.ExampleRequest{
		Title: "Hello Server",
	})
	if err != nil {
		log.Fatalf("Failed to call ExampleForever: %v", err)
	}

	// 只需要接收服务端的流数据
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("bye...")
				break
			}
			log.Printf("Failed to receive message: %v", err)
			return
		}
		fmt.Printf("[ExampleForever]receive server msg: %s\n", msg.Reply)
	}

}
