package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	protoc "Go/cmd/OldPackage/grpc_test/proto"
)

// 业务逻辑
type Server struct{}

func (s *Server) SayHello(ctc context.Context, request *protoc.HelloRequest) (*protoc.HelloReply, error) {
	return &protoc.HelloReply{
		Message: "hello, " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	// 注册server
	protoc.RegisterGreeterServer(g, &Server{}) // 实例化
	// 启动
	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis) // 没有直接结束连接
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
