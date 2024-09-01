package add

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 这里移动要用Test开头
func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d,actual %d", 4, re)
	}
}

// 跳过一些比较耗时的单元测试
func TestAdd2(t *testing.T) {
	fmt.Println("yes1")
	if testing.Short() {
		t.Skip("short 模式下跳过")
	}
	fmt.Println("yes2") // 没有运行
	re := add(1, 5)
	if re != 6 {
		t.Errorf("expect %d,actual %d", 6, re)
	}
}

// 基于表格驱动测试
func TestAdd3(t *testing.T) {
	// 匿名结构体
	// 测试数据较多的时候，把下面这样的理解为表格，快速对大量数据进行测试
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{12, 12, 24},
		{-9, 8, -1},
		{0, 0, 0},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect %d,actual %d", value.out, re)
		}
	}
}

// 性能测试
func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 456
	c = 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d+%d,expect %d,actual %d", a, b, c, actual)
		}
	}
}

const NUMBERS = 10000

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < NUMBERS; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}

	b.StopTimer()
}
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < NUMBERS; j++ {
			str = str + strconv.Itoa(j)
		}
	}

	b.StopTimer()
}
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < NUMBERS; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}

	b.StopTimer()
}
