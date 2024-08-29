package main

import (
	"flag"
	"fmt"
)

// func visit(list []int, f func(int)) {
// 	for _, v := range list {
// 		f(v)
// 	}
// }

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	// 使用匿名函数打印切片内容
	// visit([]int{1, 2, 3, 4}, func(v int) {
	// 	fmt.Println(v)
	// })

	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chichen fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
