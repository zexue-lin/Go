package main

func main() {
	//在这个例子中，iota 可以被用作枚举值：
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)

	const (
		a = iota // a = 0
		b        // b = 1
		c        // c = 2
		d = 5    // d = 5
		e        // e = 5
	)
	// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
	const (
		Apple, Banana     = iota + 1, iota + 2 // Apple=1 Banana=2
		Cherimoya, Durian                      // Cherimoya=2 Durian=3
		Elderberry, Fig                        // Elderberry=3, Fig=4

	)

	// 使用 iota 结合 位运算 表示资源状态的使用案例
	const (
		Open    = 1 << iota // 0001
		Close               // 0010
		Pending             // 0100
	)
}
