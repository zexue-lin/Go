package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if age > 10 {
		fmt.Println("大于10岁啦！")
	}

}
