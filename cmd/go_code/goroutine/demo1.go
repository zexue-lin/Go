package main

import (
	"fmt"
	"time"
)

// 主协程
func main() {

	// 主死从随

	// 1.闭包 2.for循环问题，for循环的时候 每个面两会重用
	// for循环的时候，i变量会被重用

	//匿名函数启动goroutine (每次打印的值不一样)
	for i := 0; i < 100; i++ {
		// 第一种解决办法 tem := i
		go func(i int) { // 第二种方法，调用i，使用值传递，更优雅
			fmt.Println(i)
		}(i) // 值传递
	}

	fmt.Println("main goroutine")
	time.Sleep(9 * time.Second)
}
