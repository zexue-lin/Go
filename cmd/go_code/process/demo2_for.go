package main
import "fmt"

func fordemo() {

	//	与多数语言不同的是，go中循环只支持for关键字，而不支持while，和do-while结构。
	sum := 0
	for i := 0; i <= 10; i++ {
		if sum < 10 {
			fmt.Printf("%d,", sum)
		} else {
			fmt.Println(sum)
		}
		sum += 1
	}

	// 这里类似于 do-while
	summ := 0
	for {
		summ++
		if summ > 11 {
			fmt.Println(summ)
			break
		}
	}

	fmt.Println("-----多层循环 跳出外层循环-----")
	// go的for循环和c语言的一样，区别就是go不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量，如:
Iloop: // 标签必须在首行
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				fmt.Printf("\n")
				break Iloop // 跳出外层循环
			}
			fmt.Printf("%d ", j)
		}
	}
	// 上述代码i和j的作用域被限制在各自的循环内，一旦循环结束，这些变量就超出其作用域，之后的代码将无法访问它们

	fmt.Println("-----for 中的初始化语句-----")
	// for中的初始化语句
	// 这里的step被放在for循环外面初始化，此时的step作用域比在for内初始化要大，循环结束还可以访问
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
	fmt.Println("循环结束后，还可以访问step,step=", step)

	fmt.Println("-----for 中的条件表达式-----")
	// for每次循环都会计算条件表达式是否正确，如果为true，则继续执行，
	// 条件表达式可以被忽略，忽略条件表达式后默认形成 !!! 无限循环 !!!。
	var i int
	for ; ; i++ {
		if i > 5 {
			break
		}
	}

	/*说明:
	55行，初始值未设置，两个分号之间的条件表达式未设置，会一直循环下去
	每次循环结束前都会调用结束语句 i++
	56行 判断i>5时，通过break语句跳出for循环到第60行代码

	还可以简化成下面的这样
	for忽略所以语句，此时for无限循环
	将 j++ 从 for 的结束语句放置到函数体的末尾是等效的，这样编写的代码更具有可读性。
	*/

	var j int
	for {  // 通常要配合break使用
		if j > 5 {  // 循环条件
			break
		}
		j++
	}

	/*
			氦油糕熟 还可以简化成下面这样，（前提是只有一个循环条件）
		将之前使用if i>5{}判断的表达式进行取反，变为判断 t 小于等于 5 时持续进行循环。

		下面这段代码其实类似于其他编程语言中的 while，在 while 后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。
	*/
	var t int
	for t <= 5 {   // 循环条件
		fmt.Printf("%d ", t)
		t++
	}
}
