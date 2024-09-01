package main

import "fmt"

// 接口嵌套

type MyWriter interface {
	Write(string)
}

type MyReader interface {
	Read() string
}

type MyReadWriter interface {
	MyWriter
	MyReader
	ReadWriter()
}

type SreadWriter struct {
}

func (s *SreadWriter) Write(s2 string) {
	//TODO implement me
	fmt.Println("write")
}

func (s *SreadWriter) Read() string {
	//TODO implement me
	fmt.Println("read")
	return ""
}

func (s *SreadWriter) ReadWriter() {
	//TODO implement me
	fmt.Println("read write")
}

func main() {
	var mrw MyReadWriter = &SreadWriter{}
	mrw.Read()
	// 上面是值对象接收者的情况：SreadWriter 【对于下面是通用的】
	// 有一个小细节，如果这里实例赋值的时候使用SreadWriter{}或&SreadWriter{}  都是ok的

	// 上面是指针接收者的情况: *SreadWriter
	// 下面必须要用 &SreadWriter{}   否则报错
}
