package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{}

func main() {
	s := MyStruct{}
	v := reflect.TypeOf(s)
	fmt.Printf("Type: %v Kind: %v Name: %v\n", v, v.Kind(), v.Name())
	// Type: main.MyStruct Kind: struct Name: MyStruct

	// kind返回底层类型就是struct，name顾名思义就是它的名字
}
