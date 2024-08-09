package main

import "fmt"

// 声明全局变量
var a float32 = 3.14

func main() {
	//声明局部变量
	var a int = 3
	fmt.Printf("a = %d\n", a) // 输出3

	age := 25
	name := "Alice"

	// 使用 fmt.Printf 进行格式化输出
	fmt.Printf("年龄：%d，名字：%s\n", age, name)

	// 使用 fmt.Println 进行简单输出
	fmt.Println("这是简单输出，没有格式化")
	fmt.Println(age, name)
}
