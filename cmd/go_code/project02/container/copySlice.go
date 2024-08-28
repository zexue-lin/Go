package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前面三个元素到slice2中
	fmt.Println(slice2)  // [1 2 3]

	copy(slice1, slice2) // 只会复制slice2的三个元素到slice1的前三个位置
	fmt.Println(slice1)  // [5 4 3 4 5]

	const elementCount = 1000

	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData

	copyData := make([]int, elementCount)
	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0]) // 999

	fmt.Println(copyData[0], copyData[elementCount-1]) // 0 999

	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6]) // srcData[4:6] = 4 5
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i]) // 4 5 2 3 4,前面两个已经被替换
	}

	// ---------------从切片中删除元素---------------
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a[1:])
	// 不移动数据指针，将后面的数据向开头移动
	fmt.Println(append(a[:0], a[1:]...))
	// 使用copy函数来删除开头的1个元素
	fmt.Println(a[:copy(a, a[1:])])

	seq := []string{"a", "b", "c", "d", "e"}

	index := 2 // 要删除的位置
	fmt.Println(seq[:index], seq[index+1:])
	// 追加切片的时候要用... 如果追加的是已知元素，可以不使用...
	fmt.Println(append(seq[:index], seq[index+1:]...))


}
