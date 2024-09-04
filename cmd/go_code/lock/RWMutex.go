package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁

/*
锁本质是将并行的代码串行化，使用lock会影响性能
即使设计锁，也应该保证并行

读写锁：
假设这里有两组协程，一组写数据，另一组读数据。在web场景中，读多写少
虽然有多个goroutine，但是仔细分析会发现，读协程之间的goroutine可以并发，读和写之间应该串行（写的时候不可以读）
写和写之间也不应该并行
*/

func main() {
	//var num int
	var rwlock sync.RWMutex // 表示读写锁
	var wg sync.WaitGroup
	// 有三个方法 RLock Lock Unlock

	wg.Add(6)

	// 写的goroutine

	go func() {
		time.Sleep(time.Second * 3)
		defer wg.Done()
		rwlock.Lock() // 写锁会防止别的写操作和读操作获取
		defer rwlock.Unlock()
		//num = 12
		fmt.Println("get Write Lock")
		time.Sleep(time.Second * 5)

	}()

	// time.Sleep(time.Second) // 将下面的goroutine休眠一秒
	// 读的goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock() // 读锁,读锁不会组织别人的读操作
				//fmt.Println(num)
				fmt.Println("get read lock")
				time.Sleep(500 * time.Millisecond)
				rwlock.RUnlock()
			}
		}()
	}
	// 执行顺序，1注册wg.Done 2获取读锁，3注册RUnlock，4打印num，5释放读锁，6最后执行wg.Done()
	wg.Wait()
}

//这里设置一个读操作一个写操作，先让读操作执行，3秒后拿到写锁，等待写锁5秒（此时不能进行读操作），5秒后再次进行读操作
