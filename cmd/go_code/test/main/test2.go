package main

import "fmt"

var a string

func main() {
	a = "G"
	print(a)
	f1()

	type BitFlag int
	const (
		Active  BitFlag = 1 << iota // 1 << 0 == 1
		Send                        // 1 << 1 == 2
		Receive                     // 1 << 2 == 4
	)
	fmt.Println()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
