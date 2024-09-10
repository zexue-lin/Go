package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1.建立连接  Dial => 拨号
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}

	var reply string // string有默认值
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "Tom", &reply)
	//client.Hello("jerry", &reply) 怎么做到这样呢？ 需要自己封装
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

}
