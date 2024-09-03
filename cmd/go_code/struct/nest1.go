package main

import "fmt"

type Person struct {
	name string
	age  int
}

// p Person 是接收器；print 是方法名 ;有一个返回值（类似函数的）这里省略到了
// 接收器有两种 （值传递和引用传递）
// p Person  => 值作为接收者，通过拷贝接受Person的值并只输出Person的内容
// p *Person => 接受一个指向Person的指针，并改变Person内部的值
func (p *Person) print() {
	p.age = 23
	fmt.Printf("name:%s age:%d\n", p.name, p.age)
}
func main() {
	p := Person{
		"tom", 18,
	}
	p.print()
	fmt.Println(p.age)
	/*
	 p Person  =>
	 输出：name:tom age:23
	 18

	 p *Person  =>
	 输出：name:tom age:23
	 23  改变了Person内部age的值
	*/
}
