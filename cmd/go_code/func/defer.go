package main

import "fmt"

func main() {
	// 连接数据、打开文件、开始锁、 无论如何最后都要记得关闭数据库、关闭文件、解锁

	// defer后面的代码会在函数return之前执行
	fmt.Println(deferFunc()) // 11
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("main")
	return
	// 打印 main 2 1
	// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）

}

// defer可以修改函数返回值，因为defer会放在函数return之前就执行了
func deferFunc() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}
