package main

import "fmt"

// 定义了两个接口
type MyWriter interface {
	Write(string) error
}
type Mycloser interface {
	Close() error
}

// 具体的实现（可以不写，主要关心的是方法）
// 一个结构体实现多个接口，writerCloser实现了MyWriter和Mycloser两个接口
type writerCloser struct {
	MyWriter // interface也是一个类型 ， 像放入一个写文件的实现和一个写数据库的实现
}

// 一个接口被多个结构体实现 这个结构体也可以实现上面的接口，多对多的关系
//type writer struct{}

type fileWriter struct {
	filePath string
}

func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil
}

// ----
type databaseWriter struct {
	host string
	port string
	db   string
}

func (dw *databaseWriter) Write(string) error {
	fmt.Println("writ string to database")
	return nil
}

//----

func (wc *writerCloser) Close() error {
	fmt.Println("close writerCloser")
	return nil
}

func main() {
	//var mw MyWriter = &writerCloser{}
	//var mc Mycloser = &writerCloser{}
	// 实例化
	var mw MyWriter = &writerCloser{
		&databaseWriter{}, // 实现
	}
	mw.Write("tom")
}
