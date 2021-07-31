package main

import (
	"fmt"
)

func difference(a int, b int, c int, d int) int {
	return (a * b) - (c * d)
}

func main() {
	var a, b, c, d, result int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)

	result = difference(a, b, c, d)

	fmt.Println("DIFERENCA =", result)
}
