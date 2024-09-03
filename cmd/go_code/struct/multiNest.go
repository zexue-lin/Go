package main

import "fmt"

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct {
}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {

	cp := new(CameraPhone)
	fmt.Println("我们的新手机有多种功能")
	fmt.Println("它有拍照功能:", cp.TakeAPicture())
	fmt.Println("还有打电话功能:", cp.Call())
	/*
		我们的新手机有多种功能
		它有拍照功能: Click
		还有打电话功能: Ring Ring
	*/
}
