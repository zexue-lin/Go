package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机数字
	n := rand.Intn(100)
	fmt.Println(n)

	rand.Seed(time.Now().Unix())
	fmt.Println(time.Now().Unix()) // 输出时间戳

}

