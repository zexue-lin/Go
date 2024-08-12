package main

import (
	"fmt"
	_ "strconv" // 暂时没有使用到这个包，又不想删掉，可以使用 _ 省略掉
)

// 从指针获取指针指向的值
func main() {
	println("+++++从指针获取指针指向的值+++++")
	var home string = "guangzhou baiyun" // 创建一个字符串并赋值
	var ptr *string = &home              // 对字符串取地址 ，将地址保存到ptr
	/*
		ptr 指针变量
		ptr类型是 *string
		&home home的地址
		var ptr *string = &home 指针变量ptr就指向了home的地址，
		通过*ptr就可以访问到home的值
	*/
	fmt.Printf("ptr type is: %T\n", ptr) // 打印ptr的类型

	fmt.Printf("ptr'address is: %p\n", ptr) // 打印ptr的指针地址，每次运行都会发生变化，

	var value string = *ptr                  // 对ptr指针变量取值，
	fmt.Printf("value type is: %T\n", value) // 打印取值后value的类型

	fmt.Printf("value is : %s\n", value) // 打印value的值

	var ptt *int                           // 空指针
	fmt.Printf("ptt address is:%p\n", ptt) // 默认地址：0x0  默认值：nil

	var i = 999
	var pty *int = &i
	println(&pty)                          // pty自己的地址
	println(&i)                            // i的地址
	fmt.Printf("pty address is:%p\n", pty) // pty指向的地址

	println("----使用指针修改值----")
	var num int = 9
	fmt.Printf("num address is:%p\n", &num)

	var ptr1 *int = &num
	*ptr1 = 10 // 这里修改的就是num的值，*ptr1 就直接操作num的值
	fmt.Printf("num value is:%d\n", num)

	//var num2 int = 300
	//var ptr2 *float32 = &num2

	var a int = 300
	var b int = 400

	var ppp *int = &a
	*ppp = 100
	ppp = &b
	*ppp = 200
	fmt.Printf("a=%d, b=%d,*ppp=%d", a, b, *ppp)

}
