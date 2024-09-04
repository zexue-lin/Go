package main

import (
	"fmt"
	"time"
)

// channel 要初始化
//var done = make(chan struct{}) // 空结构体

func g1(ch chan struct{}) {
	time.Sleep(1 * time.Second)

	ch <- struct{}{} // 前面是空结构体定义 后面的{}是实例化

}

func g2(ch chan struct{}) {
	time.Sleep(2 * time.Second)

	ch <- struct{}{}
}
func main() {

	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)
	//g1Channel <- struct{}{}
	//g2Channel <- struct{}{}
	go g1(g1Channel)
	go g2(g2Channel)

	//<-done // 阻塞的方法

	// 我要监控多个channel，任何一个channel返回值我都知道  用select
	// 1.某一个分支就绪了，就执行该分支
	// 2.如果两个都就绪了，先执行哪个？【重要】  答：是随机的，目的是防止饥饿（每次都执行一个的情况）
	// 应用场景:设置超时时间，给每一个分支设置时间，超时就执行超时的分支

	// 设置超时时间
	timer := time.NewTimer(5 * time.Second) // 声明一个timer对象，里面有个channel
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C: // C代表channel，1秒之后就可以访问了
			fmt.Println("timeout")
			return

		}
		time.Sleep(10 * time.Millisecond)

		// 虽然解决了 但大多时候并不会多个goroutine写同一个channel
		// 如何监控多个channel呢？

	}
}
