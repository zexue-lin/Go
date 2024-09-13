package main

import (
	"Go/cmd/OldPackage/stream_grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
	for {
		a, err := res.Recv() //socket编程的模式在使用 send recv 两个函数
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}
}
