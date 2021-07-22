package main

import "fmt"

func main() {
	var input [5]int

	fmt.Scanf("%d", &input[0])
	fmt.Scanf("%d", &input[1])
	fmt.Scanf("%d", &input[2])
	fmt.Scanf("%d", &input[3])
	fmt.Scanf("%d", &input[4])

	result := calculate(input)

	fmt.Printf("%d valores pares\n", result)
}

func calculate(input [5]int) int {
	counter := 0

	for _, number := range input {
		if !(number%2 == 0) {
			continue
		}

		counter++
	}

	return counter
}
