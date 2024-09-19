package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader_switch *bufio.Reader

func main() {
	inputReader_switch = bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字:")
	input, err := inputReader_switch.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Println("您的名字是:", input)

	switch input {
	case "Philip\r\n": // \r回车  \n换行
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	/*

	 */
}
