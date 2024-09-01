package main

import "fmt"

// 结构体嵌套
type Person struct {
	name string
	age  int

	// 下面用的较少
	//address struct {
	//	city string
	//}
}

type Student struct {
	// --------------------第一种嵌套方式--------------------
	//p Person
	//score float32

	// --------------------第二种嵌套方式，匿名嵌套--------------------
	Person
	score float32
	name  string
}

// --------------------结构体定义方法--------------------
// p Person => 接收器;   print => 方法名;  有一个返回值（类似函数的）这里省略到了
// 接收器有两种形态,值传递和引用传递（p *Person）
func (p Person) print() {
	// 1.如果函数中想要改变age的值，只能使用引用传递
	// 2.如果person对象很大，值传递要copy，影响性能，数据较大的时候使用引用传递
	p.age = 19
	fmt.Printf("name:%s,age:%d\n", p.name, p.age)
}

func main() {
	// --------------------第一种嵌套方式--------------------
	//第一种赋值方式 这里要手动输入name age score这三个字段名字
	//s := Student{
	//	p: Person{
	//		name: "bobby", age: 18,
	//	},
	//	score: 98.7,
	//}
	//// 取值
	//fmt.Println(s.p.name)

	// 第二种赋值方式
	//s := Student{}
	//s.p.name = "bobby"
	//fmt.Println(s.p.name)

	// --------------------第二种嵌套方式--------------------
	// 不需要输入字段名字
	s := Student{
		Person{
			"bobbys", 23,
		},
		97.7,
		"外部Student的name",
	}
	//// 取值 有两个name 一个是Person的，一个是Student的，有限访问外部的那个
	//fmt.Println(s.name)
	//
	//// 这里设置的是外面的那个name
	//s.name = "outside"
	//fmt.Printf("s内部的实际存储:%v", s)

	// --------------------结构体定义方法--------------------
	p := Person{
		"tom", 18,
	}
	p.print()          // 调用方式
	fmt.Println(p.age) // p Person => age = 18 ; p *Person => age = 19
	s.print()          // 调用方式

}
