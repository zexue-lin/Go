package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	po := &Person{
		name: "bobby",
	}
	po.name = "tom"
	fmt.Println(po)

	var a int = 10
	b := &a
	fmt.Println(b)

	var c *int
	fmt.Println(c) // nil

	//var p *Person// 不行 不能直接访问
	//var ps Person // 可以  但是什么都不打印
	//fmt.Println(ps.name)

	ps := &Person{} // 第一种指针初始化方式

	var emptyPerson Person
	pi := &emptyPerson // 第二种指针初始化方式

	var pp = new(Person) // 第三种初始化方式
	fmt.Println(pp.name)

	/*
		初始化的两个关键字
		make关键字： map、channel、slice的初始化
		new关键字（函数）：指针的初始化，指针不初始化会出现 nil pointer
		map必须要初始化
	*/
}
