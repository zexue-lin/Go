package main

import (
	_ "flag"
	"fmt"
	"strings"
	"unicode/utf8"
)

// 定义命令行参数
// var mode = flag.String("mode","","process mode")

func main() {

	// 解析命令行参数
	// flag.Parse()

	// fmt.Println(*mode)

	str := new(string)
	*str = "Go语言指针"
	fmt.Println(*str)

	type Weekday int
	const (
		sunday Weekday = iota
		monday
		tuesday
		wednesday
	)

	fmt.Println(wednesday)

	tip1 := "genji is a ninja"
	fmt.Println(len(tip1)) // 16
	tip2 := "忍者"
	fmt.Println(len(tip2)) // 6

	fmt.Println(utf8.RuneCountInString("忍者"))          // 2
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!")) // 11

	trace := "死神来了，死神bye bye"

	comma := strings.Index(trace, "，")
	pos := strings.Index(trace[comma:], "死神")
	fmt.Println(comma, pos, trace[comma+pos:],trace[15:])
}
