package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	quantity := 0
	input := []float64{0, 0}

	fmt.Scanf("%d", &quantity)

	for index := quantity; index >= 1; index-- {
		fmt.Scanf("%f %f", &input[0], &input[1])

		if input[1] == 0 {
			fmt.Println("divisao impossivel")

			continue
		}

		calculate(input)
	}
}

func calculate(input []float64) {
	fmt.Printf("%.1f\n", input[0]/input[1])
}
