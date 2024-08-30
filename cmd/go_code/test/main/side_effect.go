package main

import "fmt"

func main() {
	n := 0
	ptr := &n
	Multiply(5, 10, ptr)
	// 传递指针给函数可以节省内存，（因为没有复制变量的值）
	// 而且可以使函数可以直接修改外部变量的值
	// 所以不用return返回
	fmt.Println(*ptr) // 50
	fmt.Println(n)    // 50
}

func Multiply(a, b int, ptr *int) {
	*ptr = a * b
}
