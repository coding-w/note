package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"note/grpc/stream_client/hello"
	"strings"
)

type server struct {
	hello.UnimplementedHelloServiceServer
}

func (s *server) Hello(stream hello.HelloService_HelloServer) error {
	var names []string
	for {
		// 从客户端流中接收数据
		req, err := stream.Recv()
		if err == io.EOF {
			// 当客户端流结束时，构造响应并返回
			return stream.SendAndClose(&hello.HelloResponse{
				Reply: "你好, " + strings.Join(names, " ") + "!",
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", req.GetName())
		names = append(names, req.Name)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
