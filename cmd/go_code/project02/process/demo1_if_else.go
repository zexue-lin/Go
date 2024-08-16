package main
import "fmt"

func demoifelse() {

	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if age > 10 {
		fmt.Println("大于10岁啦！")
	}

	var n = 23
	if n > 11 {
		fmt.Println("test.....ok")
	} else {
		fmt.Println("no ok")
	}

	// ------------------特殊写法------------------
	// err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。
	// err != nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。
	if err := Connect(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("连接成功")
	}

}

// Connect 尝试建立连接，如果没有错就返回nil
// 返回类型为 error
func Connect() error {
	// 如果连接成功，返回nil
	return nil
}
