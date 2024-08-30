package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机数

type Rope string

func main() {
	for i := 0; i < 5; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	for i := 0; i < 5; i++ {
		a := rand.Intn(16) // 括号里可以是8 16 32 64
		fmt.Println(a)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
	fmt.Println(time.Now().Unix()) // 输出当前时间戳

	var ccd Rope = "kny"
	fmt.Printf("ccd type is%T,ccd=%v", ccd, ccd) // ccd type ismain.Rope,ccd=kny
}
