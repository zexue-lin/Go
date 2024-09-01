package main

import "fmt"

// 接口的定义
type Duck interface {
	// 方法的声明
	Gaga()
	Walk()
	Swimming()
}

// 接口的实现
type pskDuck struct {
	legs int
}

// 这里必须对应接口里面的三个方法
func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}
func (pd *pskDuck) Walk() {
	fmt.Println("this is pskduck walking")
}
func (pd *pskDuck) Swimming() {
	fmt.Println("this is pskduck swimming")
}

func main() {
	// go语言的接口 ，鸭子类型 php python
	// go语言中处处都是interface，到处都是鸭子类型 duck typing

	/*
		当看到一只鸟走起来像鸭子、游泳起来像是鸭子、叫起来也是鸭子、那么这只鸟就是鸭子
		动词 ，方法， 具备的某些方法
	*/
	// 鸭子类型强调的是食物的外部行为，而不是内部的结构，使用起来像是动态语言

	var d Duck = &pskDuck{}
	d.Walk()
}
