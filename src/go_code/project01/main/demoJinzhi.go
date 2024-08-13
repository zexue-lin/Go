package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 将二进制数 110001100 转换为十进制
	var shi = 1001101111                   // 二进制数
	var binaryStr = fmt.Sprintf("%d", shi) // 先转成字符串
	decimalValue1, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("二进制转换出错:", err)
		return
	}
	fmt.Printf("二进制 %s 转换为十进制为: %d\n", binaryStr, decimalValue1)
	fmt.Printf("decimalValue1的数据类型是: %T\n", decimalValue1)

	// 将八进制数 02456 转换为十进制
	// 下划线表示空标识符，这里用来忽略错误
	octalStr := "02456"
	decimalValue2, _ := strconv.ParseInt(octalStr, 8, 64)
	fmt.Printf("八进制 %s 转换为十进制为: %d\n", octalStr, decimalValue2)

	// 将十六进制数 0xA45 转换为十进制
	hexStr := "0xA45"
	decimalValue3, _ := strconv.ParseInt(hexStr, 0, 64)
	fmt.Printf("十六进制 %s 转换为十进制为: %d\n", hexStr, decimalValue3)

	fmt.Println("-----十进制转其他-----")

	var num = -6
	binary2 := strconv.FormatInt(int64(num), 2) // 十进制转二进制   1111011
	fmt.Printf("binary的数据类型是%T,binary=%v\n", binary2, binary2)

	binary8 := strconv.FormatInt(int64(num), 8) // 十进制转八进制
	fmt.Printf("binary的数据类型是%T,binary=%v\n", binary8, binary8)

	binary16 := strconv.FormatInt(int64(num), 16) // 十进制转十六进制
	fmt.Printf("binary的数据类型是%T,binary=%v\n", binary16, binary16)
}
