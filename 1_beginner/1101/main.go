package main

import (
	"fmt"
	"sort"
)

func main() {
	input := readInput()

	printResult(input)
}

func readInput() []string {
	input := [2]int{0, 0}
	numbers := make([]string, 0)

	for 0 <= 1 {
		fmt.Scanf("%d %d", &input[0], &input[1])

		if input[0] <= 0 || input[1] <= 0 {
			return numbers
		}

		sort.Ints(input[:])

		numbers = append(numbers, calculate(input))
	}

	return numbers
}

func calculate(input [2]int) string {
	var result string
	sum := 0

	for index := input[0]; index <= input[1]; index++ {
		if index != input[1] {
			sum += index
			result += fmt.Sprintf("%d ", index)

			continue
		}

		sum += index
		result += fmt.Sprintf("%d Sum=%d", index, sum)
	}

	return result
}

func printResult(results []string) {
	for _, result := range results {
		fmt.Println(result)
	}
}
