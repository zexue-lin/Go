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

	// 虽然解决了 但大多时候并不会多个goroutine写同一个channel
	// 如何监控多个channel呢？ 用select

	for { // 无限循环，让事件处理逻辑持续运行。监听多个通道会在任意一个case上接收到并执行
		select {
		case <-g1Channel:
			fmt.Println("g1 done") // g1快一点
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C: // C代表channel，1秒之后就可以访问了
			fmt.Println("timeout")
			return
		}
		time.Sleep(10 * time.Millisecond)
		// 作用是在 select 语句没有匹配到任何可接收的通道时，让当前 goroutine 睡眠一小段时间。这是为了避免 CPU 的空转，
		// 如果 select 语句没有匹配任何 case，它会立即再次检查通道是否准备好。这种情况下，如果没有 time.Sleep，循环可能会消耗大量的 CPU 资源。
	}
}
