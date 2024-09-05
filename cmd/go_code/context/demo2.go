package main

import (
	"fmt"
	"sync"
	"time"
)

// 虽然channel可以实现主给子routine传递消息。go还设计了一个context来实现(更优雅)，看WithCancel

var wg sync.WaitGroup

// 1.主的gorouting（main）向子goroutine传递信息使用channel方式
//var stop = make(chan struct{})

// 2.但是使用参数传递更合理
func cpuInfo(stop chan struct{}) { // 子goroutine
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("CPU info")
		}

	}
}

func main() { // 父goroutine

	// 3.作为参数，所以把这一句移到下面来了
	var stop = make(chan struct{})
	wg.Add(1)
	go cpuInfo(stop)

	time.Sleep(time.Second * 5)
	stop <- struct{}{}
	wg.Wait()
	fmt.Println("监控完成")
}
