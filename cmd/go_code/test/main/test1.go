package main

import (
	"fmt"
	_ "os"
	"runtime"
	"strings"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Print("The operating system is ", goos) // The operating system is windows

	// path := os.Getenv("PATH")
	// fmt.Print("Path is",path)
	var str1 = "hello"
	var str2 = "0"
	fmt.Println(strings.Contains(str1, str2))

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 当i=5的时候直接进入下一次循环，执行i++  直接i=6
		}
		print(i)
		print(" ")
	}
LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}
