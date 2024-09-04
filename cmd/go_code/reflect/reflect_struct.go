package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

// 为上面的结构体实现了String方法
func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// 空接口 interface{} 可以保存任何类型的值，因为所有类型都实现了空接口。
// 在这里，把一个 NotknownType 实例赋值给 secret 变量。
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	// alternative:
	// typ := value.Type()  // main.NotknownType
	fmt.Println(typ) // main.NotknownType
	knd := value.Kind()
	fmt.Println(knd) // struct

	// 循环访问结构体的字段
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}

	// 调用第一个方法, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
	//value通过reflect.ValueOf(secret)获取的反射值，它代表了NotknownType的实例
	//Method(0)获取实例的第一个方法，这里只有一个方法String()

	// .Call(nil): .Call 方法用于调用反射获取的方法。
	// nil表示该方法没有参数
	// Call 方法返回一个 []reflect.Value 类型的切片，包含该方法的返回值。
}
