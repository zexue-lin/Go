package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 共享变量 【需加锁保护】
var stop bool

// 两个地方操作stop，加一个读写锁
var rwlock sync.RWMutex

func cpuInfo() {
	defer wg.Done()
	for {
		rwlock.RLock()
		if stop {
			//不用defer: 这种方式可以避免在循环中使用 defer，确保锁在每次循环结束后立即释放。
			// 这里不用defer延迟，判断完状态直接释放锁。 可能发生资源泄漏，在 'for' 循环中调用 'defer
			// 因为 defer 在每次循环中都会产生开销。
			rwlock.RUnlock() // 在跳出循环之前先释放读锁
			break            // 跳出循环for循环，后面的语句不会被执行
		}
		rwlock.RUnlock() // 正常情况下释放读锁
		time.Sleep(2 * time.Second)
		fmt.Println("CPU info")
	}

	/*
		条件检查：进入循环后，首先调用 rwlock.RLock() 获取读锁，然后检查 stop 变量的值。
		如果 stop 为 true：
		执行 rwlock.RUnlock() 释放读锁。
		然后执行 break，跳出循环，循环体中的剩余代码（包括第二个 rwlock.RUnlock()）不会被执行。
		如果 stop 为 false：
		代码继续执行 rwlock.RUnlock() 释放读锁，然后继续循环。
		【关键点】
		break 语句的作用：break 会立即终止循环，控制权跳到循环之外，不会执行 break 后面的代码。
		因此，即使在 stop 为 true 的情况下，有两个 RUnlock() 操作，但实际上只有第一个 RUnlock() 会被执行，
		第二个 RUnlock() 不会被执行。
	*/
}
func main() {

	// 1.有一个goroutine监控cpu的信息
	// 2.我可以主动退出监控程序。解决办法1 共享变量
	wg.Add(1)
	go cpuInfo()

	time.Sleep(6 * time.Second)
	rwlock.Lock()
	stop = true
	rwlock.Unlock() // 这里不需要defer，直接释放读锁

	wg.Wait()

	fmt.Println("监控完成")
}
