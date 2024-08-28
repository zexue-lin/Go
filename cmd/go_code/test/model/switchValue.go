package model

import "fmt"

// 交换两个变量的值，不使用中间变量
func main() {
	var a = 10
	var b = 20
	a = a + b
	b = a - b // b=a+b-b  => b=a
	a = a - b // a=a-(a-b) => a = a-a+b  => a=b

	fmt.Println(a, b)
}
