package main

import "fmt"

func main() {

	// var sum int16
	// var max int16 = 100
	// var i int16 =1
	// for  ; i <= max; i++ {
	// 	if i%9 == 0 {
	// 		sum += i
	// 	}
	// }
	// fmt.Println(sum)

	// 先算一个班的平均分
	// 再计算三个班的 j
	// var classNum int = 2
	// var stuNum int = 3
	// var totalSum float64 = 0.0
	// var passCount = 0
	// for j := 1; j <= classNum; j++ {
	// 	sum := 0.0
	// 	for i := 1; i <= stuNum; i++ {
	// 		var score float64
	// 		fmt.Printf("请输入第%d班 第%d个学生的成绩: \n", j, i)
	// 		fmt.Scanln(&score)
	// 		sum += score

	// 		if score >= 60 {
	// 			passCount++
	// 		}
	// 	}
	// 	fmt.Printf("第%d个班级的平均分为:%v\n", j, sum/float64(stuNum))
	// 	totalSum += sum

	// }
	// fmt.Printf("总成绩=%v,总平均分为%v\n,及格人数为%v,", totalSum, totalSum/(float64(stuNum)*float64(classNum)), passCount)

	// -----------------金字塔-----------------
	// 一半的时候：第一层一个 第二层2个，层数等于星星数量
	// *
	// **
	// ***

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// 一整个,,规律规律规律规律规律
	//   *     2 * 层数 -1  ，空格数=总层数-当前层
	//  ***    2 * 层数 -1  ，
	// *****   2 * 层数 -1  ，
	for a := 0; a < 5; a++ {
		for k := 0; k < 5-a; k++ {
			fmt.Print(" ")
		}

		for b := 0; b < 2*a-1; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
