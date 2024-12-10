package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"note/grpc/stream_server/example"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := example.NewExampleServiceClient(conn)

	// 发送请求，接收服务端流的响应
	req := &example.ExampleRequest{Name: "Alice", Nums: int32(5)}
	stream, err := client.ExampleStream(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error: %v", err)
		}
		log.Printf("response: %v", response)

	}

}
