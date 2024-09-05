package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		// 如何做到先打印数字，此处应该等待打印另一个goroutine通知”该我打印数字了“
		<-number // 这里表示true
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter
		if i >= len(str) {
			return
		} // 防止死循环
		fmt.Printf(str[i : i+2])
		i += 2
		number <- true
	}
}
func main() {

	// 题目，使用两个goroutine交替打印序列，一个打印数字，另一个打印字母，最终效果如下：
	// 12AB34CD56EF78GH910IJ......

	go printNum()
	go printLetter()
	number <- true

	time.Sleep(90 * time.Second)
}
