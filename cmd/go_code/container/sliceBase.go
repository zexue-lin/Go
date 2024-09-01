package main

import (
	"fmt"
	"strconv"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 5; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

func main() {
	// go的slice在函数传递的时候是值传递还是引用传递
	// 其实是值传递，但是呈现出引用的效果（不完全是）

	// 其实是值传递
	couse := []string{"go", "grpc", "gin"}
	printSlice(couse)
	fmt.Println(couse)
}
