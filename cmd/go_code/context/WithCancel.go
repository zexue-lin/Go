package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 虽然channel可以实现主给子routine传递消息，go还设计了一个context来实现(更优雅)，

var wg sync.WaitGroup

// 这里使用context，关键字 ctx
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // 这本身也是一个channel，阻塞的
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("CPU info")
		}

	}
}

func main() {

	wg.Add(1)

	// 如果你的goroutine，函数中，希望被控制。（如：超时，传值）但是我不想影响原来的接口信息
	// 函数参数中的第一个参数就尽量要加上一个ctx，它是一个interface接口
	// 把这个放在第一 ctx context.Context， 后面可以再加参数
	ctx1, cancel1 := context.WithCancel(context.Background()) // 父cancel
	ctx2, _ := context.WithCancel(ctx1)                       // 子cancel

	// 可以一直派生，具有传递性

	go cpuInfo(ctx2)

	time.Sleep(time.Second * 5)

	cancel1() // 这里一旦调用cancel方法，上面的ctx.Done立马可以收到
	wg.Wait()
	fmt.Println("监控完成")
}
