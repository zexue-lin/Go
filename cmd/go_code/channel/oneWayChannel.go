package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	// in <- 1  // 报错，因为这里只读
	for num := range in {
		fmt.Printf("num=%d\n", num)
	}
}

func main() {
	// 单向channel的第一种定义方式
	//var ch1 chan int       // 双向的channel
	//var ch2 chan<- float64 // 单向channel，只能写入float64的数据
	//var ch3 <-chan int     // 单向，但是只能读

	//c := make(chan int, 3)
	//var send chan<- int = c // send-only
	//var read <-chan int = c // recv-only
	//
	//send <- 1
	//<-read

	c := make(chan int) // 双向channel
	go producer(c)

	go consumer(c)

	time.Sleep(10 * time.Second)
}
