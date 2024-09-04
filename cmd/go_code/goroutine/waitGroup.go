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

	// 这里有三个方法
	// 这里表示要监控多少个goroutine执行结束 ，这里是100个
	wg.Add(100)
	for i := 0; i < 100; i++ {

		go func(i int) {
			//一般先写这一句
			defer wg.Done() // 因为是defer，这一句肯定会在匿名函数执行完成之后才执行
			fmt.Println(i)

		}(i)
	}

	// 等到
	wg.Wait()
	fmt.Println("all done")

	// waitgroup主要用于goroutine的执行等待，Add方法要和Done成对出现
}
