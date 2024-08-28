package main

import "fmt"

func main() {
	res := fbn(3)

	fmt.Println("res=",res)
	fmt.Println("res=",fbn(4))
	fmt.Println("res=",fbn(5))
	fmt.Println("res=",fbn(6))

}

func fbn(n int) int {
	// 斐波那契数列 1 1 2 3 5 8 13
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
