package main

import "fmt"

func main() {
	cc := [3]int{1, 2, 3}
	for k, v := range cc {
		fmt.Printf("k=%d,v=%d\n", k, v)
	}

	var array1 [3]int = [3]int{3, 4}
	for i, j := range array1 {
		fmt.Printf("i=%d,j=%d", i, j)
	}

	// 数组的长度由初始化的值来计算
	array := [...]string{"apple", "banana", "cat", "dog", "egg", "food", "green", "hand", "ice", "just", "kill", "lemon", "mom", "new", "open", "post",
		"question", "red", "system", "top", "us", "view", "window", "x", "yellow", "zero"}

	for _, v := range array {
		fmt.Println(v)
	}

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"

	// d := [3]int{1, 2}
	// fmt.Println(a == d) // 编译错误无法比较 [2]int and [3]int
}
