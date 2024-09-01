package main

import (
	"Go/cmd/go_code/test/model"
	"fmt"
)

func main() {
	fmt.Println(model.HeroName)

	// 如果都是整数进行运算，则会省略小数部分
	var n float32 = 10 / 4
	println(int(n))

	// 若想得到小数部分，则需要有浮点数参与运算
	var n1 float32 = 10.0 / 4
	fmt.Printf("%.1f\n", n1)

	// 取模运算
	// 公式：a % b = a - a / b * b
	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3) // -10 - (-10) / 3 * 3 = -10 - (-3)*3 = -10 - (-9) = -10 + 9 = -1

	var date = 97
	var week = date / 7
	var day = date % 7
	fmt.Printf("还有%d星期%d天放假\n", week, day)

	var hua float32 = 134.2
	var tmp float32 = 5.0 / 9 * (hua - 100)
	fmt.Printf("对应的摄氏温度是%v\n", tmp)

	println("-------逻辑短路运算--------")
	var i int = 10
	if i > 9 && test() {
		println("ok")
	}

}
func test() bool {
	fmt.Println("test....")
	return true
}
