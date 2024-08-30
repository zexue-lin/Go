package main

import (
	"fmt"
	"strings"
)

func main() {
	// ------------字符串包含关系------------
	var str1 = "hello"
	var str2 = "lo"
	// strings.Contains() 用于检查字符串 str1 中是否包含字符串 str2。
	fmt.Println(strings.Contains(str1, str2)) // true

	// ------------判断子字符串或字符在父字符串中出现的位置（索引）------------
	str3 := "woshinidwosiswoswuubuwose"
	str4 := "s"
	fmt.Println(strings.Index(str3, str4))     // 输出2，判断str4在字符串str3中的索引位置
	fmt.Println(strings.LastIndex(str3, str4)) // 输出2，判断str4在字符串str3中的最后的索引位置

	str5 := "我是你爹的你的你的"
	var str6 string = "你"

	fmt.Println(strings.Index(str5, str6)) // 输出6 因为一个中文占三个字节

	// 如果需要查询非 ASCII 编码的字符在父字符串中的位置，使用以下函数来对字符进行定位：
	//fmt.Println(strings.IndexRune(str5, str6)) // 输出6 因为一个中文占三个字节

	// ------------字符串替换------------
	strPep := strings.Replace(str3, "wos", "xxn", 2) // 最多替换两个wos
	fmt.Println(strPep)
	strPep1 := strings.Replace(str3, "wos", "xxn", -1) // 替换全部的wos
	fmt.Println(strPep1)

	// ------------统计字符串出现的字数------------
	fmt.Println(strings.Count(str3, str4)) // 出现了五次s

	// ------------重复字符串------------
	fmt.Println(strings.Repeat(str5, 2)) // 重复两次str5

	// ------------修改字符串大小写------------
	var str7 string = "aBcDeFhIjK"
	fmt.Println(strings.ToLower(str7))
	fmt.Println(strings.ToUpper(str7))

	// ------------修剪字符串------------
	var str8 string = " aBcDeourFhIjK "
	var str9 string = " aBcDeourFhIjK "
	fmt.Println(strings.TrimSpace(str8))

	// 也可以剔除指定的字符串
	fmt.Println(strings.TrimLeft(str9, " "))
	// 剔除字符串左边的空格和aBc TrimRight()是右边的
	fmt.Println(strings.TrimLeft(str9, " aBc"))

	// ------------分割字符串------------
	strr := "The quick brown fox jumps over the lazy dog"
	strr1 := strings.Fields(strr) // 返回切片
	fmt.Println(strr1)
	// 遍历切片，使用-连接
	for _, val := range strr1 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	st1 := "Go|The ABC of Go|25"
	st2 := strings.Split(st1, "|")
	fmt.Println(st2)
	for _, val := range st2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	st3 := strings.Join(st2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", st3)
}
