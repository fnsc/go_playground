package main

import "fmt"

func calculate(input [6]float64) int {
	var counter int = 0
	for _, number := range input {
		if number >= 0 {
			counter++
		}
	}

	return counter
}

func main() {
	var input [6]float64

	fmt.Scanf("%f", &input[0])
	fmt.Scanf("%f", &input[1])
	fmt.Scanf("%f", &input[2])
	fmt.Scanf("%f", &input[3])
	fmt.Scanf("%f", &input[4])
	fmt.Scanf("%f", &input[5])

	result := calculate(input)

	fmt.Printf("%d valores positivos\n", result)
}
