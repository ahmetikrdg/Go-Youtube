package main

import "fmt"

func main() {
	fmt.Println(plus(3, 5))

	res := plusplus(5, 5, 5)
	fmt.Println(res)
}

func plus(a int, b int) int {
	return a + b
}

func plusplus(a int, b int, c int) int {
	return a + b + c
}
