package main

import "fmt"

func main() {
	var address string = "北京长城123\nasdsa"
	fmt.Println(address)

	// 字符串一旦赋值，不能修改
	// 反引号使用
	var str = `
var address string = "北京长城123\nasdsa"
fmt.Println(address)
`
	fmt.Println(str)

	var str1 = "hello" + "world"
	str1 += "hahah"
	fmt.Println(str1) // helloworldhahah

	// 当一个拼接的操作很长时，可以分行写,加号要留在上面,每行的结尾必须时加号
	var str2 = "hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world" + "hello" + "world"
	fmt.Println(str2)

	fmt.Println("---------")
	// 基本数据类型的默认值
	// %v是按照变量的值输出
	var a int
	var b float32
	var c float64
	var d byte
	var isMarried bool
	var name string
	fmt.Printf("a=%d,b=%f,c=%f,d=%c,isMarried=%t,name=%v", a, b, c, d, isMarried, name)
}
