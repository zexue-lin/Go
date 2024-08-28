package main

import "fmt"

func main() {
	// var a = [3]int{1, 2, 3}
	// fmt.Println(a,a[1:2]) // [1 2 3] [2]

	// ---------------------1.从指定范围生成切片---------------------
	fmt.Println("1.从指定范围生成切片")
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}

	// 区间
	fmt.Println(highRiseBuilding[10:15])

	// 索引20到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])

	fmt.Println(highRiseBuilding[:2])

	// ---------------------2.表示原有的切片---------------------
	// 生成切片的格式中，当开始和结束位置都被忽略时，生成的切片将表示和原切片一致的切片，
	// 并且生成的切片与原切片在数据内容上也是一致的，代码如下：
	fmt.Println("2.表示原有的切片")
	a := []int{1, 2, 3}
	fmt.Println(a[:])

	// ---------------------3.直接声明新的切片---------------------
	fmt.Println("3.直接声明新的切片")
	var strList []string
	var numList []int
	var numEmptyList = []int{} // 此时的 numListEmpty 已经被分配了内存，只是还没有元素。
	fmt.Println(strList, numList, numEmptyList)
	fmt.Println(len(strList), len(numList), len(numEmptyList))

	fmt.Println("判定为空:")
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numEmptyList == nil) // 此时的 numListEmpty 已经被分配了内存，只是还没有元素。

	// ---------------------4.使用make()函数构造切片---------------------
	fmt.Println("4.使用make()函数构造切片:")
	// 第一个参数指切片元素类型，第二个：分配多少元素；第三个预分配元素数量，不影响第二个参数，只为解决性能问题
	c := make([]int, 2)
	d := make([]int, 20, 30) // 内部存储空间已经分配了 30 个，但实际使用了 20 个元素。

	fmt.Println(c, d)
	fmt.Println(len(c), len(d))

	// ---------------------5.使用append()为切片添加元素---------------------
	fmt.Println("5.使用append()为切片添加元素:")
	var e []int
	e = append(e, 1)
	e = append(e, 1, 2, 3)
	e = append(e, []int{4, 5, 6}...) // 追加一个切片, 切片需要解包
	fmt.Println(e) // [1 1 2 3 4 5 6]

	// 切片在扩容时，容量cap的扩展规律是按容量的2倍进行扩充，如下:
	var numbers []int
	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	// 除了在切片的尾部追加，我们还可以在切片的开头添加元素：
	fmt.Println("在切片的开头添加元素:")
	e = append([]int{0}, e...)
	e = append([]int{-3,-2,-1},e...)
	fmt.Println(e) // [-3 -2 -1 0 1 1 2 3 4 5 6]
}
