package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
)

// rpc远程过程调用，如何做到像本地调用
type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int { //
	// 传输协议 http
	// 数据格式协议 url里的
	req := HttpRequest.NewRequest()
	res, err := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	if err != nil {
		log.Fatalf("Error making request: %s", err)
	}

	body, err := res.Body()
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	fmt.Println(string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	fmt.Println(Add(2, 2))
}
