package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var os = runtime.GOOS // 查看当前运行的操作系统类型
	fmt.Println(os)

}
