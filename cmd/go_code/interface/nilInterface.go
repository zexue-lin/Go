package main

import "fmt"

var any interface{}

func main() {

	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = false
	fmt.Println(any)

}
