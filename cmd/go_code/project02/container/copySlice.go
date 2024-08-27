package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前面三个元素到slice2中
	fmt.Println(slice2)  // [1 2 3]

	copy(slice1, slice2) // 只会复制slice2的三个元素到slice1的前三个位置
	fmt.Println(slice1)  // [5 4 3 4 5]

	const elementCount  = 1000
}
