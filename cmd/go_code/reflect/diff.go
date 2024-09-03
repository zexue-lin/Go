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

	i := 42
	vInt := reflect.TypeOf(i)
	fmt.Printf("Type: %v Kind: %v Name: %v\n", vInt, vInt.Kind(), vInt.Name())
}
