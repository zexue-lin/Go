package main

import (
	"fmt"
	"reflect"
)

// 反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
func reflectFn(i interface{}) {
	// 使用 reflect.TypeOf(i)获取到 类型对象v
	v := reflect.TypeOf(i)    // 返回类型
	typ := v.Kind()           // 底层种类（类型种类）
	name := v.Name()          // 返回类型名字
	val := reflect.ValueOf(i) // 返回值
	fmt.Printf("Type:%v Kind:%v Value:%v Name:%v\n ", v, typ, val, name)
}

func main() {
	a := "stt"
	b := 1
	reflectFn(a)
	reflectFn(b)

	type MyInt int
	var m MyInt = 5
	v := reflect.TypeOf(m)
	vv := v.Kind() // Kind() 总是返回底层类型：
	mm := v.Name()
	va := reflect.ValueOf(m)
	fmt.Printf("MyInt:%v Type is:%v Name:%v Value:%v\n", m, vv, mm, va)

}

/*
在 Go 语言中，reflect.TypeOf(i)、v.Kind()和v.Name()虽然有一定联系，但它们的用途和返回值有以下区别：
reflect.TypeOf(i)：
返回一个reflect.Type类型的值，表示参数i的类型信息。这个类型包含了很多关于该类型的详细信息，比如名称、包路径、是否是指针类型、是否是接口类型等。它提供了全面的类型描述。
例如，对于一个自定义结构体类型，它可以告诉你结构体的名称、字段信息等。

v.Kind()：
这里v是通过reflect.TypeOf(i)获取到的类型对象。Kind方法返回一个reflect.Kind类型的值，表示该类型的底层种类。
reflect.Kind是一个枚举类型，包括Bool、Int、String、Struct、Ptr等，表示基本的数据类型种类。它提供了一种更简洁的方式来快速确定一个值的基本类型类别，而不需要了解完整的类型信息。
例如，对于一个整数类型，无论它是int、int8、int16等具体的整数类型，Kind都会返回reflect.Int。

v.Name()：
同样，v是通过reflect.TypeOf(i)获取到的类型对象。Name方法返回这个类型的名称。
如果是命名类型（例如自定义的结构体、接口、类型别名等有明确名称的类型），它会返回这个类型的名称。如果不是命名类型，比如内置类型（int、string等）、未命名结构体等，那么它会返回一个空字符串。
*/
