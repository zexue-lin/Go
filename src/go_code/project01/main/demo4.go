package main

// Go浮点数和转义字符 查看变量的数据类型和大小
import (
	"fmt"
	"unsafe"
)

func main() {
	// \r回车，从当前行的最前面开始输出，覆盖掉之前的内容
	fmt.Println("你的牛爷爷\r我是")                        //这里的我是把你的替换掉了
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河南\t北京") // \t制表符  \n换行符

	//该区域的数据值可以在同一类型范围内不断变化
	var i = 100 //定义并初始化变量i的值为100
	i = 200
	i = 300
	fmt.Println(i)                  //值为最后一次赋值 300
	fmt.Printf("i 的数据类型为 %T \n", i) // 查看变量的数据类型
	// 查看变量占用的字节大小和整数类型
	fmt.Printf("i的类型为%T,i占用的字节数是 %d \n", i, unsafe.Sizeof(i))

	num := 5.1234e2
	num1 := 5.1234e2
	num2 := 5.1234e-2
	fmt.Println("num=", num, "num1=", num1, "num2=", num2) // num= 512.34 num1= 512.34 num2= 0.051234

}
