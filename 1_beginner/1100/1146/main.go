package main

import (
	"fmt"
	"strings"
)

func main() {
	input := 0
	control := true

	for control {
		fmt.Scanf("%d", &input)

		if input == 0 {
			control = false

			break
		}

		printResult(input)
	}
}

func printResult(input int) {
	var result string

	for index := 1; index <= input; index++ {
		result += fmt.Sprintf("%d ", index)
	}

	result = strings.TrimRight(result, " ")

	fmt.Println(result)
}
