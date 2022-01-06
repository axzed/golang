package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外不能用:=简短格式
// 不是全局变量, 是包内部变量
var (
	aa = 3
	ss = "kk"
	bb = "薛文朝"
)

// 变量定义
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d    %q\n", a, s)
}

// 变量赋初值 变量一旦用到一定要用到(go语言非常严格)
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 变量推断类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 变量定义简易写法 推荐用这种格式
func variableShorter() {
	// 只能函数内使用
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 复数
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

// 强制类型转换
// go语言对这一点拿的特别死
func triangle() {
	var a, b int = 3, 4
	var c int
	// 以为c是int型所以要显式转换, Sqrt要用float64来算
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 包内常量, go语言内常量不需要全部大写以为go语言大小写有规矩
const filename = "abc.txt"

func consts() {

	// 这里没有指定a, b的具体类型所以下面Sqrt中不需要强制类型转换
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举类型
func enums() {
	// 1.iota常量自动生成器，每个一行，自动加一
	// 2.iota常量赋值使用
	// 3.iota遇到const,重置为0
	// 4.可以只写一个iota
	// 5.同一行内iota值是一样的
	const (
		cpp = iota // 表明这组常量是自增值的
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	triangle()
	euler()
	consts()
	enums()
}
