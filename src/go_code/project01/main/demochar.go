package main

// Go字符类型
import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1, "c2=", c2) // 按照Unicode码输出 输出的是对应的 Unicode编码值

	// 如果希望输出字符。需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	var c3 rune = '北'
	fmt.Printf("c3=%c\n", c3)

	var sr rune = '你'
	var st string = "你的我的"
	fmt.Printf("sr=%c st=%s\n", sr, st)

}
