package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// bufio 包提供的缓冲读取器 (buffered reader) 来读取数据
var inputReader *bufio.Reader

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入一些信息:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	} else {
		//fmt.Printf("您输入的内容是:%s", input) // 包含一个换行符

		fmt.Printf("您输入的内容是:%s", strings.TrimSpace(input)) // 没有换行符
	}
}

/*
inputReader 是一个指向 bufio.Reader 的指针
main里面第一行代码会创建一个读取器
ReadString 返回读取到的字符串，我们会一直读取键盘输入，直到回车键 (\n) 被按下。

得到的input变量后面会包含结尾的换行符\n

strings.TrimSpace() 来去除多余的空白字符。
*/
