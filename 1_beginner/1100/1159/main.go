package main

import "fmt"

func main() {
	input := 1

	for input != 0 {
		fmt.Scanf("%d", &input)

		if input == 0 {
			return
		}

		if isOdd(&input) {
			input++
		}

		fmt.Println(*calculate(&input))
	}
}

func isOdd(input *int) bool {
	return *input%2 != 0
}

func calculate(input *int) *int {
	i := 1
	result := 0

	for i <= 5 {
		result += *input

		*input += 2
		i++
	}

	return &result
}
