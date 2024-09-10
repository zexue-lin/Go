package main

import "fmt"

func add(a, b int) int {
	res := a + b
	return res
}

func main() {
	fmt.Println(add(1, 2))
}
