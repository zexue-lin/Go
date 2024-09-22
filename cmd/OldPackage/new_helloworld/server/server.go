package main

import (
	"Go/cmd/OldPackage/new_helloworld/handler"
	"Go/cmd/OldPackage/new_helloworld/server_proxy"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 1.实例化一个server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	// 3.启动服务.for循环防止断开
	for {
		conn, _ := listener.Accept() // 当一个新的连接进来的时候，
		go rpc.ServeConn(conn)
	}

}
