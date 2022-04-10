package main

import "fmt"

func main() {
	fmt.Println("hello go")

	s2 := make([]int, 2, 4);
	s2 = append(s2, 1, 3)

	fmt.Printf("%v", s2)
}
