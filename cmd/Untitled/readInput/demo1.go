package main

import "fmt"

var (
	firstName, lastName string
	i                   int
	f                   float32
)

/*
Scanln() 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
Scanf() 与其类似，除了 Scanf() 的第一个参数用作格式字符串，用来决定如何读取。
*/
func main() {
	fmt.Println("请输入你的全名:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s", firstName, lastName)

}
