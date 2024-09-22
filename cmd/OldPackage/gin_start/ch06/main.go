package main

import (
	"Go/cmd/OldPackage/gin_start/ch06/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson" // jsonpb 包就不用了 已经过时了
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtobuf", returnProto)

	router.Run(":8083")
}

func returnProto(c *gin.Context) {
	course := []string{"python", "go", "微服务"}
	user := &proto.Teacher{
		Name:   "jerry",
		Course: course,
	}

	// 将 Protobuf 转换为 JSON 字符串
	jsonData, err := protojson.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法转换为 JSON"})
		return
	}

	// 返回 JSON 数据
	c.Data(http.StatusOK, "application/json", jsonData)
}

/*
1.Protobuf 二进制数据：cmd/OldPackage/helloworld/proto/main/main.go
这里的反解proto.Unmarshal 被用来将二进制 Protobuf 数据 rsp 反序列化为 HelloRequest 结构体；
这种做法常见于 gRPC 服务间通信，Protobuf 数据直接传输并解码，效率很高。

2.JSON 格式 Protobuf 数据：此文件
如果你传输的是 JSON 格式的 Protobuf 对象（通常是为了与前端系统交互），你可以使用 protojson.Unmarshal 来解析

对比：
1.Protobuf 二进制数据：使用 proto.Marshal 和 proto.Unmarshal 进行序列化和反序列化，效率更高，适合服务间通信或高效数据传输场景。
2.JSON 格式 Protobuf 数据：使用 protojson.Marshal 和 protojson.Unmarshal，适合需要将 Protobuf 数据以 JSON 格式传输、或者前后端交互需要 JSON 格式的场景。
*/
func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"usr"`
		Message string
		Number  int
	}
	msg.Name = "jerry"
	msg.Message = "这是一个测试json"
	msg.Number = 20

	c.JSON(200, msg)
	// http.StatusOK 等效于 200
}
