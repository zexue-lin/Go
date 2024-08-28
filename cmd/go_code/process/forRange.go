package main

import "fmt"

func main() {

	// -----------------字符串遍历--传统方式-----------------
	// var str = "hello,world!"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i]) // 数组下标
	// }

	// -----------------字符串遍历--forRange-----------------
	// var str1 = "golang"
	// for key,val:=range str1{
	// 	fmt.Printf("key=%d,val=%c\n",key,val)
	// }

	// -----------------字符串遍历(含有中文)--传统方式-----------------
	// var str2 string = "hello,world!广东"
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("%c\n", str2[i]) // 数组下标
	// }

	// 如果字符串中含有中文，传统的方式会出错，后面中文部分乱码
	// 因为是按照字节来统计的，而一个汉字在utf8中对应3个字节

	// 解决办法---> 把str2转换为切片
	// var str2 string = "hello,world!广东"
	// str3 := []rune(str2)  // 把数组转成切片 []rune
	// for i := 0; i < len(str3); i++ {
	// 	fmt.Printf("%c\n", str3[i]) // 数组下标
	// }

	// -----------------字符串遍历(含有中文)--forRange-----------------
	var str1 = "agolang广东"
	for key, val := range str1 {
		fmt.Printf("key=%d,val=%c\n", key, val)
	}

	fmt.Printf("%c", str1[0]) // 输出a
	/* 结果如下:
	key=0,val=g
	key=1,val=o
	key=2,val=l
	key=3,val=a
	key=4,val=n
	key=5,val=g
	key=6,val=广
	key=9,val=东
	因为"广"占用3个字节 占用了 6 7 8 这三个字节。
	*/
}
