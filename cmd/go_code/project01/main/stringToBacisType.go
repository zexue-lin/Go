package main

import (
	"fmt"
	"strconv"
)

// string转成基本数据类型
func main() {

	var str string = "true"
	var b bool
	/* b, _ = strconv.ParseBool(str)
	说明
	1.strconv.ParseBool(str)函数会返回两个值，(value bool,err error)
	2.我只想接受vale bool，使用 _ 忽略掉 err
	*/

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T b=%v\n", b, b)

	var str2 string = "12345"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type is %T, n1 is %v\n", n1, n1)
	fmt.Printf("n2 type is %T, n2 is %v\n", n2, n2)

	var str3 = "123.353"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type is %T,f1 is %v\n", f1, f1)

	var str4 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type is %T, n3 is %v\n", n3, n3) // 转成功了但是值为0

}
