package main

import "fmt"

func main() {

	var i int = 100
	// 希望将 i => float32
	var n float32 = float32(i)
	var str byte = byte(i)
	fmt.Printf("i=%v, n=%.2f,str=%c,", i, n, str) // i=100, n=100.00,str=d
	// n str:被转换的是变量存储的值，变量本身(i)的数据类型并没有发生变化
	// n str:被转换的是变量存储的值，变量本身(i)的数据类型并没有发生变化
	// n str:被转换的是变量存储的值，变量本身(i)的数据类型并没有发生变化
	fmt.Printf("i的数据类型是%T\n", i) // int

	// 在转换过程中，如果将 int64 转换成 int8，范围【-128~127】，不会报错
	// 只是转换的结果是按照溢出处理，
	var c int64 = 1234
	var d int8 = int8(c)
	println(d)

	// 测试题目
	var n1 int32 = 12
	var n2 int64
	var n3 int8
	// 下面错误写法
	//n2 = n1 + 20 // 无法将 'n1 + 20' (类型 int32) 用作类型 int64
	//n3 = n1 + 20 // 无法将 'n1 + 20' (类型 int32) 用作类型 int8
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println(n2, n3)

}
