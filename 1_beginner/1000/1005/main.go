package main

import (
	"fmt"
)

func average(grade1 float64, grade2 float64) float64 {
	const WEIGHT1 = 3.5
	const WEIGHT2 = 7.5

	return ((grade1 * WEIGHT1) + (grade2 * WEIGHT2)) / (WEIGHT1 + WEIGHT2)
}

func main() {
	var grade1, grade2, result float64

	fmt.Scanf("%f", &grade1)
	fmt.Scanf("%f", &grade2)

	result = average(grade1, grade2)

	fmt.Printf("MEDIA = %.5f\n", result)
}
