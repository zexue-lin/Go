package main

import "fmt"

// 结构体定义
// 定义描述一个人的结构体
type Person struct {
	name    string
	age     int
	address string
	height  float64
}

func main() {
	// ------------type关键字：定义类型别名------------
	// 别名实际上是为了更好地理解代码
	// 实际上MyInt指向了int，本质上是一样的
	// 编译的时候，类型别名会被直接替换为原本的类型
	type MyInt = int // 类型别名
	var i MyInt
	fmt.Printf("i type is %T\n", i)  // i type is int

	// ------------类型定义 自定义类型(这里没有等于号)------------
	type MyInt1 int // 自定义的类型，基于已有的类型自定义一个类型
	var j MyInt1 = 12
	fmt.Printf("j type is %T\n", j)  // j type is main.MyInt1

	// ------------扩展 interface------------
	// 接口的类型判断
	var a interface{} = "abc"
	switch a.(type) { // 拿到a的数据类型
	case string:
		fmt.Println("字符串")
	}
	// 这里是直接 断言 直接把a的数据类型转换成string
	//m := a.(string)

	// ------------初始化结构体------------
	// 也就是类型的集合

	// 初始化结构体
	// 第一种所有参数都要赋值 第二种(键值对的方式)只需要赋值想要填写的字段（第二种更灵活）
	p1 := Person{"bobby1", 18, "广州市", 1.8}
	p2 := Person{
		name: "bobby2",
		//age:  18,
		address: "北京",
		//height:  1.78,
	}

	var persons []Person
	persons = append(persons, p1, p2)
	persons = append(persons, Person{
		name: "bobby3",
		age:  20,
	})
	fmt.Println(persons)
	// 第三种
	person2 := []Person{
		// 省略的方式,不指明字段
		{"bobby1", 18, "广州市", 1.8},
		// 指明字段的方式
		{
			age: 34,
		},
		{
			name: "bobbyi",
		},
	}
	fmt.Println(person2) // [{bobby1 18 广州市 1.8} { 34  0} {bobbyi 0  0}]
	//第四种

	var p Person
	p.age = 20
	fmt.Println(p.age)  // 20
	fmt.Println(p.height) // 未初始化默认值为0

	per2 := new(Person)
	per2.age = 23
	per2.height= 175
	fmt.Println(per2.age)  
	fmt.Println(per2.height)
	fmt.Printf("per2 type is:%T\n",per2) // per2 type is:*main.Person
	
	// ------------匿名结构体------------
	fmt.Println("匿名结构体")

	// 第一个花括号是结构体的定义
	addres := struct {
		province string
		city     string
		address  string
	}{
		// 这里是对上面定义结构体的实例化
		province: "北京市",
		city:     "通州区",
		address:  "xxx街道",
	}
	fmt.Println(addres.city)
	// 这样的好处就是一次性使用，在函数内部使用，代码封装性好
}
