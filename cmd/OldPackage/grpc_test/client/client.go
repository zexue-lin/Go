package main

import (
	protoc "Go/cmd/OldPackage/grpc_test/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// 调用server.go

func main() {
	// 拨号连接
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 远程调用
	c := protoc.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &protoc.HelloRequest{Name: "Jerry"}) // 指针类型
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
