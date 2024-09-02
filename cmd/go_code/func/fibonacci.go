// package main

// import "fmt"

// func main() {
// 	res := 0
// 	for i := 1; i < 10; i++ {
// 		res = fbn(i)
// 		fmt.Printf("fbn(%d)=%d\n",i,res)
// 	}
// }

//	func fbn(n int) (res int) {
//		// 斐波那契数列 1 1 2 3 5 8 13
//		if n <= 2 {
//			res = 1
//		} else {
//			res = fbn(n-1) + fbn(n-2)
//		}
//		return
//	}
//
// fibonacci_memoization.go
package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

/*
Output: LIM=40:
normal (fibonacci.go): the calculation took this amount of time: 4.730270 s
     with memoization: the calculation took this amount of time: 0.001000 s
*/
