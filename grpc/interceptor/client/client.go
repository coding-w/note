package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"note/grpc/interceptor/example"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("conn is err %v", err)
	}
	defer conn.Close()
	client := example.NewExampleServiceClient(conn)
	response, err := client.Example(context.Background(), &example.ExampleRequest{
		Title: "123",
	})
	if err != nil {
		log.Printf("call err %v", err)
	}
	log.Printf("response %v", response)

}
