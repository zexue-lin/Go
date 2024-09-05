package main

import (
	"fmt"
)

// goroutine之间的通信

func main() {
	/*
			不要通过系统内存来通信，而要需要通信来实现内存共享

			在其他语言中，如python、php、java，多线程编程的时候，两个routine之间通信最常用的方式是一个全局变量
			也会提供消息队列的机制，python-queue java，消费者和生产者之间的关系。
			生成者生产了消息放在队列里面，消费者从队列里面消费数据。
		channel 再加上语法糖，使队列更加简单
	*/
	// 要定义类型
	var msg chan string
	//if msg == nil {
	//	fmt.Println("msg is nil")
	//}

	// 初始化
	msg = make(chan string, 1) // 这里的1是缓冲区，一个空间,设为0 就是无缓冲的channel，只要大于0就是有缓冲
	//channel初始化值如果为0，那么放值进去会阻塞，发生死锁。deadlock
	//有缓冲和无缓冲的channel
	//msg = make(chan string, 0)  // 无缓冲的channel
	//msg = make(chan string, 10) // 有缓冲的channel

	// 语法糖 其实是一个符号  <-   把右边的值放到左边
	msg <- "Tom"  // 存值到channel中
	data := <-msg // 通过变量名取值

	fmt.Println(data) // Tom

	// 如果想要解决无缓冲的channel，要启动一个goroutine
	msg = make(chan string, 0)
	go func(msg chan string) { // go有一种happen-before的机制，可以保障data := <-msg先于msg <- "Tom1"执行，可以放心消费无缓冲的数据
		data := <-msg
		fmt.Println(data)
	}(msg)
	msg <- "Tom1"
	//time.Sleep(time.Second * 2)

}
