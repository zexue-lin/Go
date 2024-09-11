package client_proxy

import (
	"Go/cmd/OldPackage/new_helloworld/handler"
	"net/rpc"
)

// 封装一个代理

type HelloServiceStub struct {
	*rpc.Client
}

// 在go语言中没有类、对象  就意味着没有初始化方法
func NewHelloServiceClient(protol, addr string) HelloServiceStub {
	conn, err := rpc.Dial(protol, addr)
	if err != nil {
		panic("connect error!")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName, request, reply)
	if err != nil {
		return err
	}
	return nil
}
