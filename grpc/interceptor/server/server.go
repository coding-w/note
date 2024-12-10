package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"note/grpc/interceptor/example"
	"time"
)

type server struct {
	example.UnimplementedExampleServiceServer
}

func (server) Example(ctx context.Context, request *example.ExampleRequest) (*example.ExampleResponse, error) {
	title := request.Title
	reply := fmt.Sprintf("receive msg:[%s],reply:%d", title, time.Now().Unix())
	return &example.ExampleResponse{Reply: reply}, nil
}

func main() {
	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor()),
	)
	example.RegisterExampleServiceServer(s, &server{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("----------------")
		fmt.Println(ctx)
		fmt.Println(req)
		fmt.Println(info)
		fmt.Println(handler)
		fmt.Println("----------------")
		return handler(ctx, req)
	}
}
