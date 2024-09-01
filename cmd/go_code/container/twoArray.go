package main

import "fmt"

func main() {
	var couInfo [3][4]string
	couInfo[0] = [4]string{"go", "1h", "b1", "初级"}
	couInfo[1] = [4]string{"java", "2h", "b2", "中级"}
	couInfo[2] = [4]string{"C", "3h", "b3", "高级"}

	fmt.Println(len(couInfo))

	for i := 0; i < len(couInfo); i++ {
		for j := 0; j < len(couInfo[i]); j++ {
			fmt.Println(couInfo[i][j] + " ")
		}
	}

	for _, row := range couInfo {
		for _, cloumn := range row {
			fmt.Print(cloumn + " ")
		}
		fmt.Println()
	}

	for _, row := range couInfo {
		fmt.Println(row)
	}
}
