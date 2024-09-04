package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var lock sync.Mutex

func g1() {
	time.Sleep(time.Second)

	lock.Lock()
	defer lock.Unlock()

	done = true

}

func g2() {
	time.Sleep(2 * time.Second)

	lock.Lock()
	defer lock.Unlock()

	done = true
}
func main() {

	// 需求:我们有两个goroutine都在执行，但是我想在主的goroutine中，当某一个执行完成以后，这个时候立马通知我
	// 这里的实现方式是全局共享变量来完成的，（但在go语言中并不推崇这么做）
	// sel2.go 使用channel来完成，使代码更优雅
	go g1()
	go g2()
	for {
		if done {
			fmt.Println("done")
			time.Sleep(10 * time.Millisecond)
			return
		}
	}
}
