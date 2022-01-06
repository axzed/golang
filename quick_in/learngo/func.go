package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) // 当你只想接收一个返回值的时候要用下划线代替另外的返回值
		return q, nil
	default:
		return 0,
			fmt.Errorf("unsuported operation:" + op)
		//panic() // 停止运行并报错
	}
}

// 带余除法
func div(a, b int) (q int, r int) {
	return a / b, a % b
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// go里没有重载这种花哨的东西, 可变参数列表
// ...int就是随便传入多少个number都可以
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(3, 4, "/"))
	fmt.Println(div(13, 3))
	// 接收两个返回值
	q, r := div(13, 5)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 10,
	))
	fmt.Println(sum(1, 2, 3, 4))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
