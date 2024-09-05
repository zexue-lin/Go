package main

import (
	"fmt"
	"sync"
)

// time.Sleep(10 * time.Second)
// 不写time.Sleep
// 子的goroutine如何通知主goroutine自己结束了，主的goroutine如何知道子goroutine已经结束了

func main() {

	var wg sync.WaitGroup

	// 这里表示要监控多少个goroutine执行结束，增加计数器
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			// 表示一个goroutine执行结束，计数器-1
			defer wg.Done() // 因为是defer，这一句肯定会在匿名函数执行完成之后才执行
			fmt.Println(i)

		}(i)
	}

	// 会发生阻塞，下面代码不会执行，等计数器归0，等待所有子 goroutines 完成。
	wg.Wait()
	fmt.Println("all done")

	// waitgroup主要用于goroutine的执行等待，Add方法要和Done成对出现
}
