package main

import "fmt"

// 接口一些细节问题
func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mPrint2(data interface{}) {
	fmt.Println(data)
}

// ------
// 打印的值不重要，类型myinfo名字不重要，有Error这个方法最重要
// 反复强调，方法方法最重要
type myinfo struct{}

func (mi *myinfo) Error() string {
	return "我不是error"
}

// ------

func main() {
	// 接口传接口可以
	var data = []interface{}{
		"Tom", 18, 1.80,
	}
	mPrint(data...)

	// 切片打散传递给接口不行
	var data2 = []string{
		"Tom", "jerry",
	} // 切片打散传递给接口 不可以
	// 迂回路线 变成这种才可以
	// 声明一个名为datai的切片，元素类型是interface
	// 遍历data2，将data2里面的值添加到datai中
	// 将一个string类型的切片转换为可以容纳任意类型的接口切片
	// 切片也分int类型切片和string类型切片等
	var datai []interface{}
	for _, value := range data2 {
		datai = append(datai, value)
	}

	mPrint(datai...)
	mPrint2(datai)

	var err error
	err = &myinfo{}
}
