package main

import (
	"fmt"
	"strings"
)

// 准备一个字符串
var str = "hello world"

func main() {

	foo := func() {
		// 匿名函数中访问str
		// 匿名函数中并没有定义str，此时，str就被引用到了匿名函数中形成了闭包
		str = "hello buddy"
		fmt.Println(str) // 输出 hello buddy
	}
	// 调用匿名函数
	foo()
	myfunc := autoAdd()
	for i := 0; i < 5; i++ {
		fmt.Println(myfunc())

	}

	// github 上的闭包例子
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	// 书写一个工厂函数而不是针对每种情况都书写一个函数。
	// 下面的函数演示了如何动态返回追加后缀的函数：
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	addBmp("file")  // returns: file.bmp
	addJpeg("file") // returns: file.jpeg
}

func autoAdd() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
