package main

import (
	"fmt"
)

func main() {
	var a, R float64
	var n float64 = 3.14159

	fmt.Scanf("%f", &R)

	a = n * (R * R)

	fmt.Printf("A=%.4f\n", a)
}
