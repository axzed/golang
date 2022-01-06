package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	temp := n
	if temp == 0 {
		return strconv.Itoa(temp)
	}
	result := ""
	// for循环可以没有初始值
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// strconv.Itoa 数字转字符串int to a
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// for变成while
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err) // panic停下报错
	}
	scanner := bufio.NewScanner(file)
	// 只有一个条件相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("xwc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 --> 1101
		convertToBin(7238),
		convertToBin(0), // 空串
	)
	printFile("abc.txt")
}
