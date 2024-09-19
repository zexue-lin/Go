package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("打开文件出错了\n" +
			"1.文件是否存在？\n" +
			"2.您是否又权限打开此文件？\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')

		fmt.Printf("读取到的文件是:%s\n", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

/*
尝试打开 input.txt 文件
如果打开成功则注册defer事件
开始读取文件并打印
遇到文件结尾跳出循环
在main函数退出之前，defer 后面的操作会被调用，确保文件被关闭

它确保了即使在处理过程中发生异常，文件也能被正确关闭，避免了资源泄露。

变量 inputFile 是 *os.File 类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）
*/
