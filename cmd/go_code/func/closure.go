package main

import "fmt"

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
}

func autoAdd() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}
