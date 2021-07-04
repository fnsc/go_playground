package main

import (
	"fmt"
)

func average(grade1 float64, grade2 float64, grade3 float64) float64 {
	const WEIGHT1 = 2
	const WEIGHT2 = 3
	const WEIGHT3 = 5

	return ((grade1 * WEIGHT1) + (grade2 * WEIGHT2) + (grade3 * WEIGHT3)) / (WEIGHT1 + WEIGHT2 + WEIGHT3)
}

func main() {
	var grade1, grade2, grade3, result float64

	fmt.Scanf("%f", &grade1)
	fmt.Scanf("%f", &grade2)
	fmt.Scanf("%f", &grade3)

	result = average(grade1, grade2, grade3)

	fmt.Printf("MEDIA = %.1f\n", result)
}
