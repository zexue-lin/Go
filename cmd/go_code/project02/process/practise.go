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
	// for a := 1; a <= 5; a++ {
	// 	for k := 0; k < 5-a; k++ {
	// 		fmt.Print(" ")
	// 	}

	// 	for b := 1; b <= 2*a-1; b++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// 打印空心金字塔
	// 每层的第一个和最后一个是*，其余的是空格
	// 最后一层全是*
	// var floor = 5
	// for a := 1; a <= floor; a++ {
	// 	for k := 1; k <= floor-a; k++ {
	// 		fmt.Print(" ")
	// 	}

	// 	for b := 1; b <= 2*a-1; b++ {
	// 		if b == 1 || b == 2*a-1 || a == floor {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}

	// 	fmt.Println()
	// }

	// 打印空心菱形
	var floor = 5
	for a := 1; a <= floor; a++ {
		for k := 1; k <= floor-a; k++ {
			fmt.Print(" ")
		}

		for b := 1; b <= 2*a-1; b++ {
			if b == 1 || b == 2*a-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for a := floor - 1; a > 0; a-- {
		for k := 1; k <= floor-a; k++ {
			fmt.Print(" ")
		}

		for b := 1; b <= 2*a-1; b++ {
			if b == 1 || b == 2*a-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// 输出9*9乘法表
	// for a := 1; a < 10; a++ {
	// 	for b := 1; b <= a; b++ {
	// 		fmt.Printf("%d*%d=%d ", b, a, a*b)
	// 	}
	// 	fmt.Println()
	// }
// var c3c [3] int = [3]int{1,2,3}

}
