package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"note/grpc/stream_bidirectional/example"
	"sync"
	"time"
)

type server struct {
	example.UnimplementedExampleServiceServer
}

func (server) Example(stream example.ExampleService_ExampleServer) error {
	var wg sync.WaitGroup
	wg.Add(2)
	// 向客户端发送消息
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			err := stream.Send(&example.ExampleResponse{
				Reply: fmt.Sprintf("[Example]Server Stream Send msg index:[%d]", i),
			})
			if err != nil {
				if err == io.EOF {
					log.Println("Client closed the connection.")
				}
				log.Fatal("stream send err:", err)
			}
			if i >= 5 {
				log.Println("stop stream send.")
				return
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				log.Println("Client closed the connection.")
				break
			}
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				break
			}
			log.Printf("[Example] Received client msg: %s\n", response.Title)
		}
	}()
	wg.Wait()
	return nil
}
func (server) ExampleForever(request *example.ExampleRequest, stream example.ExampleService_ExampleForeverServer) error {
	// 打印客户端发送的消息
	log.Printf("[ExampleForever] Received client message: %s\n", request.Title)

	// 使用 WaitGroup 等待 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(1)

	// 启动一个 goroutine 来发送消息
	go func() {
		defer wg.Done() // 确保 goroutine 完成后调用 Done

		i := 0
		for {
			i++
			resp := &example.ExampleResponse{
				Reply: fmt.Sprintf("[ExampleForever] Server Stream Send msg index: [%d]", i),
			}

			// 尝试发送消息
			if err := stream.Send(resp); err != nil {
				if err == io.EOF {
					log.Println("Client closed connection, stopping sending messages.")
					return
				}
				// 记录错误并停止发送消息
				log.Printf("Failed to send message: %v\n", err)
				return
			}
			// 当达到5条消息时停止发送
			if i >= 5 {
				log.Println("Stopping message sending after 5 messages.")
				return
			}
			// 模拟延迟
			time.Sleep(time.Second)
		}
	}()

	// 等待 goroutine 完成
	wg.Wait()

	return nil
}

func main() {
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	example.RegisterExampleServiceServer(s, &server{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	defer s.Stop()
}
