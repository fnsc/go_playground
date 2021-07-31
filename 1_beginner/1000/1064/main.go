package main

import "fmt"

func main() {
	var input [6]float64

	fmt.Scanf("%f", &input[0])
	fmt.Scanf("%f", &input[1])
	fmt.Scanf("%f", &input[2])
	fmt.Scanf("%f", &input[3])
	fmt.Scanf("%f", &input[4])
	fmt.Scanf("%f", &input[5])

	average, quantity := calculate(input)

	fmt.Printf("%d valores positivos\n%.1f\n", quantity, average)
}

func calculate(input [6]float64) (float64, int) {
	counter, sum := 0, 0.0

	for _, number := range input {
		if !(number >= 0) {
			continue
		}

		counter++
		sum += number
	}

	return sum / float64(counter), counter
}
