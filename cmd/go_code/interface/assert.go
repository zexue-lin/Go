package main

import (
	"fmt"
	"strings"
)

// 断言
// 没有解决 还是只能int类型，因为返回类型为int
//func multiadd(a, b interface{}) int {
//	ai, ok := a.(int)
//	if !ok {
//		panic("not int type")
//	}
//	bi, ok := b.(int)
//	return ai + bi
//}

func multiadd(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("Unsupported type")
	}
}

func main() {
	// 实现不同数据类型的加法
	a := 1.0
	b := 2.0
	re := multiadd("hi ", "tom")
	fmt.Println(multiadd(a, b)) // 3
	fmt.Println(re)             // hitom
	// 这里的3 和 hitom 是interface类型

	// 如果想要使用 需要先断言
	// 断言：用在处理各种不同类型的分支代码
	res, _ := re.(string)
	fmt.Println(strings.Split(res, " "))
	fmt.Println(re)

}
