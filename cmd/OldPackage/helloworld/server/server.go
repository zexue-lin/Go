package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

// Hello 第一个参数string 第二个是指针
func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "Hello" + request
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
	conn, _ := listener.Accept() // 当一个新的连接进来的时候，
	rpc.ServeConn(conn)
	/*
		一大串代码都是net好像和rpc没关系
		rpc调用中的有几个问题需要解决: 1.callid 2.序列化和反序列化
	*/
}
