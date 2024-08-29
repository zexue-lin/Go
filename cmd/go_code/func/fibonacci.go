package main

import "fmt"

func main() {
	res := 0
	for i := 1; i < 10; i++ {
		res = fbn(i)
		fmt.Printf("fbn(%d)=%d\n",i,res)
	}
}

func fbn(n int) (res int) {
	// 斐波那契数列 1 1 2 3 5 8 13
	if n <= 2 {
		res = 1
	} else {
		res = fbn(n-1) + fbn(n-2)
	}
	return
}
