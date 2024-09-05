package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// WithValue WithTimeout

// 超时功能 在运行6秒之后，自动取消
var wg sync.WaitGroup

// 这里使用context，关键字 ctx
func cpuInfo(ctx context.Context) {
	// 这里能拿到一个请求的id
	fmt.Printf("traceid:%s\r\n", ctx.Value("traceid"))
	// 记录一些日志，这次请求是哪个craceid发起的
	// 在代码维护和排查中很容易

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

	// 1.ctx1, cancel1 := context.WithCancel(context.Background()) // 父cancel
	//ctx2, _ := context.WithCancel(ctx1)                       // 子cancel

	// 2.timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), time.Second*6)

	//go cpuInfo(ctx)

	// 3.这里是手动休眠6秒后主动调用cancel，但是使用WithTimeout就不用了
	//time.Sleep(6 * time.Second)
	//cancel()

	// 4.WithDeadLine 在时间点cancel

	// 5.WithValue
	valueCtx := context.WithValue(ctx, "traceid", "gtj87i")
	go cpuInfo(valueCtx)

	wg.Wait()
	fmt.Println("监控完成")
}
