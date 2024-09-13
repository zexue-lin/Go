package main

import (
	helloworld "Go/cmd/OldPackage/helloworld/proto"

	"fmt"

	"google.golang.org/protobuf/proto" // 要用这个包才行，视频里的不能用
)

type Hello struct {
	Name string `json:"name"`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	rsp, _ := proto.Marshal(&req) // 具体的编码是如和做到的，百度一下protobuf的原理
	// fmt.Println(len(rsp))

	// 使用json做对比，传输同一个数据，name = bobby。  使用protobuf 压缩比更高
	// 通过长度比较，差不多两倍的差距
	// jsonStruct := Hello{Name: "bobby"}
	// jsonRsp, _ := json.Marshal(jsonStruct)
	// fmt.Println(len(jsonRsp))

	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)                    // 反解
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses) // 能正确解析回来
}

/*
&req: 这是一个指向 Protobuf 消息类型的指针
res: 序列化后的字节切片（[]byte），即你的消息对象被转换成了一个字节数组。
*/
