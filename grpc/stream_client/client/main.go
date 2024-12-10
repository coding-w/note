package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"note/grpc/stream_client/hello"
	"os"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)
	stream, err := client.Hello(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	// 发送多个请求到服务端
	names := []string{"Alice", "Bob", "Charlie"}
	if len(os.Args) > 1 {
		clear(names)
		names = os.Args[1:]
	}

	for _, name := range names {
		err = stream.Send(&hello.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not send request: %v", err)
		}
		time.Sleep(time.Second * 5) // 模拟延迟
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}

	log.Printf("Response from server: %s", response.Reply)

}
