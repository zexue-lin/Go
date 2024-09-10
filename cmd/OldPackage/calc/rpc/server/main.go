package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http://127.0.0.1:8000/add?a=1&b=2 比如这样写 http://127.0.0.1:8000/method=add&a=1&b=2
	// 返回的格式: josn{"data":3}
	// add的远程方法
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// 写业务逻辑
		_ = r.ParseForm()                    // 解析参数
		fmt.Println("path:", r.URL.Path)     // 请求路径
		a, _ := strconv.Atoi(r.Form["a"][0]) // i代表int，这里表示将某一个值A转换成int
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{"data": a + b})
		_, _ = w.Write(jData)
	})

	// 解决了三个问题
	/*
		1.CallID => r.URL.Path,每个url可以代表一个方法，
		2.数据的传输协议 => http的url的参数协议    【了解一下http协议的本质】
		3.网络传输协议 => http
	*/

	http.ListenAndServe(":8000", nil) // 监听8000端口
	//性能不高，而且是 http1.x 的协议
}
