package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	// switch不用写break, 除非要fallthrough
	g := ""
	switch {
	case score < 0 || score > 100:
		// panic中断执行
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	// 这个函数返回两个值 contents(文件信息) err(出错信息)
	// go中if可以像其他语言中的for一样用
	// if可以赋值出了if语句就不能用了
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grade(0),
		grade(100),
		grade(300),
	)
}
