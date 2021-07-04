package main

import (
	"fmt"
	"math"
)

func volume(radius int) float64 {
	const pi = 3.14159

	return (4.0 / 3.0) * pi * math.Pow(float64(radius), 3)
}

func main() {
	var result float64
	var radius int

	fmt.Scanf("%d", &radius)

	result = volume(radius)

	fmt.Printf("VOLUME = %.3f\n", result)
}
