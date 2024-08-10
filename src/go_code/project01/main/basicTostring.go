package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 省略类型关键词更简洁
	var num1 = 99
	var num2 = 23.456
	var b = true
	// 这里不省略是因为如果省略，编译器可能推断其数据类型为rune（int32），因为字符常量在Go中默认是rune
	var myChar byte = 'h'
	var str string

	// 第一种方法转换【推荐】 fmt.Sprintf()
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T,str=%s\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T,str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	// %q	用于以带引号的形式输出值，以便更清晰地展示字符串或字符的值。，必要时会采用安全的转义表示
	fmt.Printf("str type is %T,str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	// 方式二 strconv 包里的函数
	var num3 = 99
	var num4 = 99.99
	//var b2 = true

	str = strconv.FormatInt(int64(num3), 10) // base:10 就代表转成10进制的数，2就是2进制的数
	fmt.Printf("str type is %T,str=%v\n", str, str)
	// num4：要进行格式化的浮点数；'f'：指定格式化的样式。（这里是常规样式）；10:表示小数点后的数字位数，64:指定 num4 的类型为 float64 。
	str = strconv.FormatFloat(num4, 'f', 10, 64)

	fmt.Printf("str type is %T,str=%v\n", str, str)

	// 为什么这里的str这么多 不会报错
}
