package main

import (
	"bytes"
	"fmt"
)

func main() {

	myfunc(2, 3, 4, 5)

	myfunc2([]int{7, 8, 9}) // 不推荐

	var (
		v1 int     = 1
		v2 int64   = 678
		v3 string  = "xsw"
		v4 float32 = 1.23
	)
	myPrintf(v1, v2, v3, v4)
	/* 结果如下
	1 is an int value.
	678 is an int64 value.
	xsw is an string value.
	1.23 is an unknow value.
	*/

	// 输入三个字符串，将他们连接成一个字符串
	fmt.Println(joinStrings("pig ", "and", " rat")) // pig and rat

	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "sttr", true))
	/* 结果如下
	value:100type:int
	value:sttrtype:string
	value:truetype:bool
	*/

	print(11, 22, 33)
}

// 可变参数类型,固定变量类型
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg) // 2 3 4 5
	}
}

// 可变参数类型
func myfunc2(args []int) {
	for _, arg := range args {
		fmt.Println(arg) //  7 8 9
	}
}

// 任意类型的可变参数
func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is an string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknow value.")

		}
	}
}

// 遍历可变参数列表——获取每一个参数的值
func joinStrings(lists ...string) string {
	// 定义一个可变大小的字节缓冲区，快速地连接字符串
	// 可以高效地进行字节和字符串的写入和读取操作。
	var b bytes.Buffer

	// 遍历可变参数列表lists，类型为[]string
	for _, s := range lists {
		// 将遍历出的字符串连接写入字节数组
		b.WriteString(s) //b 中就包含了 lists 中所有字符串连接后的结果。
	}

	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

// 获得可变参数类型——获得每一个参数的类型
func printTypeValue(slist ...interface{}) string {
	// 字节缓冲区作为快速字符串连接用
	var b bytes.Buffer

	//遍历参数
	for _, s := range slist {
		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)

		// 类型的字符串描述
		var typeString string

		// 对s进行类型判断
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		// 写字符串前缀
		b.WriteString("value:")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString("type:")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}

func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	// 将变量在 print 的可变参数列表中添加...后传递给 rawPrint()。
	rawPrint(slist...) // 11 22 33
	// rawPrint("fmt",slist) // [11 22 33]
}

func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		fmt.Println(a)
	}
}
