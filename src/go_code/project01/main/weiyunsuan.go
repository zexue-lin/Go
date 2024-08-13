package main

import (
	"fmt"
)

func main() {
	var i = 2
	var j = 3

	fmt.Println(i & j) // 2
	fmt.Println(i | j) // 3
	fmt.Println(i ^ j) // 1

	var c = -1
	var d = 2
	fmt.Println(c & d) // 2
	fmt.Println(c | d) // -1
	fmt.Println(c ^ d) // -3

	// 与运算
	a := 5 // 二进制: 0101
	b := 3 // 二进制: 0011
	andResult := a & b
	fmt.Printf("与运算结果: %d\n", andResult) // 结果为 1，二进制: 0001

	// 或运算
	orResult := a | b
	fmt.Printf("或运算结果: %d\n", orResult) // 结果为 7，二进制: 0111

	// 非运算
	notA := ^a
	fmt.Printf("非运算结果: %d\n", notA) // 结果为 -6，取反后的二进制补码表示: 1010

	// 按位异或运算
	xorResult := a ^ b
	fmt.Printf("按位异或运算结果: %d\n", xorResult) // 结果为 6，二进制: 0110

	// 左移运算
	leftResult := a << 2
	fmt.Printf("左移两位运算结果: %d\n", leftResult) // 结果为 20，二进制: 10100

	// 右移运算
	rightResult := a >> 1
	fmt.Printf("右移一位运算结果: %d\n", rightResult) // 结果为 2，二进制: 0010

	// 先对第二个操作数取反，然后执行按位与操作。
	e := 6  // 二进制: 0110
	f := 11 // 二进制: 1011  取反后: 0100
	result := e &^ f
	fmt.Printf("与非运算结果: %d\n", result) // 结果为 4，二进制: 0100
}
