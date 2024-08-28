package main

import "fmt"

func main() {

	var month int
	fmt.Println("请输入一个月份:")
	fmt.Scanf("%d",&month)

	switch month{
	case 3,4,5:
		fmt.Println("spring")
	case 6,7,8:
		fmt.Println("summer")
	case 9,10,11:
		fmt.Println("autumn")
	case 12,1,2:
		fmt.Println("winter")
	}
}