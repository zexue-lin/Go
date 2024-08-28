package main

import "fmt"

func fire() {
	fmt.Println("fire")
}

func main() {
	var f func()  // 定义f变量，数据类型是func()
	f = fire
	f()  // 调用函数，实际上调用的就是fire()
}
