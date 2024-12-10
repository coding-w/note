package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"note/grpc/stream_server/example"
	"strconv"
	"time"
)

type server struct {
	example.UnimplementedExampleServiceServer
}

func (s *server) ExampleStream(req *example.ExampleRequest, stream example.ExampleService_ExampleStreamServer) error {
	log.Printf("Received ExampleStream Request %+v\n", req)
	for i := 0; i < int(req.Nums); i++ {
		response := example.ExampleResponse{Reply: "你好，" + req.Name + "! 第" + strconv.Itoa(i) + "次问候"}
		err := stream.Send(&response)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	example.RegisterExampleServiceServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
