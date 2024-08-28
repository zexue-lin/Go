package input

import "fmt"

func main() {
	//方式1
	//ftm.Scanln 按照标准输入读取数据，直到遇到换行符
	var name string
	var age int
	var sal float32
	var isPass bool
	//fmt.Println("请输入姓名")
	//// 程序执行到这里的时候，会等待用户输入
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//fmt.Println("请输入是否通过考试")
	//fmt.Scanln(&isPass)
	//
	//fmt.Printf("姓名=%v,年龄=%v,薪水=%v,是否通过考试=%t", name, age, sal, isPass)

	// 方式2
	// fmt.Scanf,可以按指定的格式输入
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名=%v,年龄=%v,薪水=%v,是否通过考试=%t", name, age, sal, isPass)
}
