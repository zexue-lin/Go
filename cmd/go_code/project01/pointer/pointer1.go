package main

import (
	_ "flag"
	"fmt"
)

// 定义命令行参数
// var mode = flag.String("mode","","process mode")

func main() {
	
	// 解析命令行参数
	// flag.Parse()

	// fmt.Println(*mode)

	str :=new(string)
	*str = "Go语言指针"
	fmt.Println(*str)

	type Weekday int
	const(
		sunday Weekday = iota
		monday
		tuesday
		wednesday
	)

	fmt.Println(wednesday)
}