package main

import (
	"fmt"
	"os"
)

// 将整个文件的内容读到一个字符串里
func main() {
	inputFile := "./input.txt"
	outputFile := "./output.txt"
	//buf, err := ioutil.ReadFile(inputFile) // 该方法第一个返回值的类型是 []byte   字节切片
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("文件内容是:%s\n", string(buf))

	// 将读取到的文件写入到另一个文件中
	err = os.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err)
	}
}

/*

fmt.Fprintf(os.Stderr, "File Error: %s\n", err) 是一个用于将错误信息输出到标准错误流（stderr）的语句。
程序通常有三个标准文件描述符：标准输入（stdin）、标准输出（stdout）和标准错误（stderr）

从 Go 1.16 开始，ioutil 包被标记为过时，并且建议使用 os 和 io 包中的函数来替代。

0644 权限允许文件所有者读写文件，而文件所属组和其他用户只能读取文件。
*/
