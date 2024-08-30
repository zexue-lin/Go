package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 30.08.2024

	t = time.Now().UTC()
	fmt.Println(t)

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week)) // 一周后的时间
	fmt.Println(week_from_now)

	// 格式化时间
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))

	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
