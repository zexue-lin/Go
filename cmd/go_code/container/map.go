package main

import (
	"fmt"
	"sync"
)

func main() {

	var courseMap = map[string]string{
		"go":   "go功臣是",
		"grpc": "grpc入门",
		"gin":  "gin入门的",
	}
	courseMap["mysql"] = "mysql原理"
	fmt.Println(courseMap)

	// 只有一个参数的时候获取的是键，也可以遍历map
	for key := range courseMap {
		fmt.Println(key, courseMap[key])
	}

	// 判断一个键是否在map中
	fmt.Println("-------------判断一个键是否在map中-------------")
	// 这里的ok是bool类型的,如果找到了php这个键，ok=true，否则ok=false
	// if!ok => if ok==false(!!!这里是判断语句!!!)
	d, ok := courseMap["php"]
	if !ok {
		fmt.Println("not in")
	} else {
		fmt.Println(d)
	}
	// 上面的写法等价于下面这个
	if _, ok := courseMap["php"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("in")
	}

	fmt.Println("-------------2.sync.Map-------------")
	var scene sync.Map

	// 将键值对保存到sync.map
	scene.Store("green", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 取键值对
	fmt.Println(scene.Load("london"))
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("k&v:", k, v)
		return true
	})
}
