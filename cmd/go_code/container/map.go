package main

import (
	"fmt"
	"sync"
)

func main() {
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
