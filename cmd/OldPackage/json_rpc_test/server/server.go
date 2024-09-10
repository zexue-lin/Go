package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

// 给结构体定义方法：Hello 第一个参数string 第二个是指针
func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "Hello, " + request
	return nil
}

func main() {
	// 1.实例化一个server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	// 2.将结构体注册到rpc当中，注册一个handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3.启动服务
	for {
		conn, _ := listener.Accept() // 当一个新的连接进来的时候，
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	// 底层数据传过来是json格式都可以解析
}
