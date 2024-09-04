package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//v.SetFloat(3.1415) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	// 报错原因 v是不可设置的，使用v.CanSet 测试它是否可设置
	fmt.Println("v可以设置嘛：", v.CanSet()) // 返回false
	// 上面v := reflect.ValueOf(x)函数通过一个x拷贝创建了一个v，v的值改变并不能更改原始的x

	// 但是传递x的地址就可以
	v = reflect.ValueOf(&x)
	fmt.Println(v.Type())                // 类型是 *float64
	fmt.Println("传地址可以设置嘛：", v.CanSet()) // 也还是false

	// 要想让其可设置我们需要使用 Elem() 函数，这间接地使用指针：v = v.Elem()
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)         // 3.4
	fmt.Println("settability of v:", v.CanSet()) // true

	v.SetFloat(3.1415)
	fmt.Println(v.Interface()) // 3.1415
	fmt.Println(v)             // 3.1415
}
