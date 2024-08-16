package main

import "fmt"

func main() {

	// 可以将很多个 if else-if else 改写成一个 switch。
	var key byte
	var a string
	var age int
	fmt.Println("请输入一个字符,一个家长英文，一个年龄")
	fmt.Scanf("%c %s %d", &key, &a, &age)

	/* ++++++++++++案例一++++++++++++ */
	// 不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行
	// switch key {
	// case 'a':
	// 	fmt.Println("周一")
	// case 'b':
	// 	fmt.Println("周二")
	// case 'c':
	// 	fmt.Println("周三")
	// default:
	// 	fmt.Println("默认周日")
	// }

	// 2.case后面可以是一个常量、变量、函数
	switch test(key) {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	default:
		fmt.Println("默认周日")
	}

	// 4.一分支多值
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	default:
		fmt.Println("不对哦")
	}

	// 7.switch 后也可以不带表达式，类似 if- else分支来使用

	switch {
	case age == 10:
		fmt.Printf("年龄是:%d", age)
	case age == 20:
		fmt.Printf("年龄是:%d", age)
	case age == 30:
		fmt.Printf("年龄是:%d", age)
	default:
		fmt.Println("默认年龄")
	}

	// case中也可以对范围进行判断
	var score int
	switch {
	case score > 90:
		fmt.Println("成绩优秀")
	case score > 60:
		fmt.Println("成绩及格")
	case score < 60:
		fmt.Println("成绩不及格")
	default:
		fmt.Println("默认成绩")
	}

	// 8.switch后也可以之间声明/定义一个变量，分号结束，【不推荐】
	switch grade := 90; {
	case grade > 90:
		fmt.Println("成绩优秀~")
	case grade > 60:
		fmt.Println("成绩及格~~")
	case grade < 60:
		fmt.Println("成绩不及格~~~")
	default:
		fmt.Println("~默认成绩~")
	}

	// 9. switch穿透 fallthrough
	var num1 = 10
	switch num1 {
	case 10:
		fmt.Println("~10~")
		fallthrough //默认穿透一层  执行了10，继续执行下一层20
	case 20:
		fmt.Println("~20~")
	case 30:
		fmt.Println("~30~")
	default:
		fmt.Println("num1 num1")
	}
}

func test(char byte) byte {
	return char + 1
}
