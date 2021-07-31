package main

import "fmt"

func main() {
	input := readInput()

	printResult(input)
}

func readInput() [][2]int {
	input := [2]int{0, 0}
	numbers := make([][2]int, 0)

	for 0 < 1 {
		fmt.Scanf("%d %d", &input[0], &input[1])

		if input[0] == input[1] {
			return numbers
		}

		numbers = append(numbers, input)
	}

	return numbers
}

func printResult(input [][2]int) {
	for _, number := range input {
		if number[0] < number[1] {
			fmt.Println("Crescente")

			continue
		}

		fmt.Println("Decrescente")
	}
}
