package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	h "note/grpc/base/hello"
)

type server struct {
	h.UnimplementedHelloServiceServer
}

// Hello 实现 grpc方法
func (s *server) Hello(ctx context.Context, in *h.HelloRequest) (*h.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &h.HelloResponse{Reply: "你好，" + in.Name + "！"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个新的 gRPC 服务器实例
	s := grpc.NewServer()
	// 将 server 实例注册到 gRPC 服务器中
	h.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	// 启动 gRPC 服务器，开始监听并处理客户端请求
	// 程序将一直运行，直到手动停止或遇到错误
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
