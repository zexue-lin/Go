package main

// 互斥锁 Mutex atomic
import (
	"fmt"
	"sync"
)

/*
锁 - 对共享资源的竞争，解决办法:将进行竞争的代码原子化
atomic 原子包  --> 将简单的+ - 原子化

锁更灵活，可以锁住一个代码段。
*/

var total int32
var wg sync.WaitGroup

var lock sync.Mutex

// 锁能复制吗？  不能！ 复制之后就失去了锁的效果
// 加了锁，相当于把这两个函数的执行串行化了
func add() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		total += 1    // 竞争，两个函数同时执行这一句，要解决使用锁就行
		lock.Unlock() // 执行完记得把锁释放掉

		// 解决办法2 用atomic包
		//atomic.AddInt32(&total, 1)
	}
}

func sub() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()

		// atomic.AddInt32(&total, -1)
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
