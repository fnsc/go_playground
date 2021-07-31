package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanf("%d", &input)

	grades := make([][3]float64, input)

	for index, grade := range grades {
		fmt.Scanf("%f %f %f", &grade[0], &grade[1], &grade[2])

		grades[index] = grade
	}

	result := calculate(grades)

	printResult(result)
}

func calculate(grades [][3]float64) []float64 {
	const WEIGHT1 = 2
	const WEIGHT2 = 3
	const WEIGHT3 = 5
	const WEIGHT_DIVISOR = (WEIGHT1 + WEIGHT2 + WEIGHT3)
	averages := make([]float64, 0)

	for _, grade := range grades {
		average := ((grade[0] * WEIGHT1) + (grade[1] * WEIGHT2) + (grade[2] * WEIGHT3)) / WEIGHT_DIVISOR
		averages = append(averages, average)
	}

	return averages
}

func printResult(result []float64) {
	for _, average := range result {
		fmt.Printf("%.1f\n", average)
	}
}
