package main

import "fmt"

func main() {
	input := 0

	read(&input)

	result := readCases(&input)

	printResult(result)
}

func read(input *int) {
	fmt.Scanf("%d", input)
}

func readCases(input *int) *[][3]int {
	cases := make([][3]int, *input)

	for i := 0; i < *input; i++ {
		read(&cases[i][0])
		read(&cases[i][1])
		calculate(&cases[i])
	}

	return &cases
}

func calculate(input *[3]int) {
	i := 1

	if isEven(&input[0]) {
		input[0]++
	}

	for i <= input[1] {
		input[2] += input[0]
		input[0] += 2

		i++
	}
}

func isEven(input *int) bool {
	return *input%2 == 0
}

func printResult(result *[][3]int) {
	for _, item := range *result {
		fmt.Println(item[2])
	}
}
