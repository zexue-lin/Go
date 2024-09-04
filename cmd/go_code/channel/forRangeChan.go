package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int        // 声明一个名为msg的整数类型通道
	msg = make(chan int, 2) // 初始化，缓冲区大小为2，意味着可以向这个通道中发送最多两个值而不会阻塞发送操作，直到有接收操作发生。

	go func(msg chan int) {
		//data := <-msg
		//fmt.Println(data) // 这里只取了一次值
		//
		//data = <-msg
		//fmt.Println(data) // 1.再取一次

		// 2.比较麻烦，使用for range
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all down") // 没有打印出来是因为阻塞了，缓冲区为2 3.使用close 就可以打印出来
	}(msg)

	msg <- 1
	msg <- 2

	//close(msg) //3.和其他编程语言有很大区别，其他编程语言不可以close一个队列；go中可以，一旦close，forRange就自动退出

	// 4.关闭之后还能不能放值进channel? 不能，已经关闭的channel不能再放值了
	//msg <- 3 // panic: send on closed channel

	// 5.先close，再取值呢，可以的
	close(msg) // 已经关闭的channel可以再取值，不能再放值了
	d := <-msg
	fmt.Println(d)
	time.Sleep(time.Second * 3)

}
