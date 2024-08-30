package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println("次数", i)
	}

	// 不使用for的循环
	i := 0
START:
	fmt.Println(i)
	i++
	if i <= 15 {
		goto START
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
