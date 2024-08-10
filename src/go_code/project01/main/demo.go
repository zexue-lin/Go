package main

// Go变量定义
import (
	"fmt"
	"sort"
)

// 使用关键词 var 和括号，可以将一组变量定义放一起
// var (
//
//	a int
//	b string
//	c []float32
//	d func() bool
//	e struct {
//		x int
//	}
//
// )

func main() {
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)
	//x := 100
	//aa, ss := 1, "abc" // 定义字符串要使用 双引号
	ch := 'G'       // 单引号用于定义 rune 类型，是go中字符类型，实际上是一个32位的整数，表示Unicode编码点，单引号内只能包含一个字符
	fmt.Println(ch) // 输出: 71 (字符 'G' 的 Unicode 编码值)

	s := IntSlice{6, 9, 5, 3, 2, 4, 1}
	sort.Sort(s)
	fmt.Println(s)

	aaa, _ := GetData()
	_, bbb := GetData()
	fmt.Println(aaa, bbb)
}
func GetData() (int, int) {
	return 100, 200
}

// IntSlice 这部分不属于一个独立的函数，而是为自定义类型 IntSlice 定义的方法。
// IntSlice 定义了一个名为IntSlice 的类型，它本质上是一个 []int 类型的别名
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }           //这是一个方法，用于返回切片的长度。
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }      //这个方法用于比较切片中索引为 i 和 j 的元素的大小关系。
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] } //用于交换切片中索引为 i 和 j 的元素的值。
