package handler

// 为了解决名称冲突问题
const HelloServiceName = "handler/HelloService"

type NewHelloService struct{}

// 给结构体定义方法：Hello 第一个参数string 第二个是指针
func (s *NewHelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "Hello, " + request
	return nil
}
