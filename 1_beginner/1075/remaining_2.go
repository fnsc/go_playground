package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanf("%d", &input)

	startingPoint := getStartingPoint(input)

	printResult(startingPoint, input)
}

func getStartingPoint(input int) int {
	step := 1
	for index := 1; index <= 10000; index++ {
		if !(index%input == 2) {
			continue
		}

		return index
	}

	return step
}

func printResult(startingPoint, input int) {
	for index := startingPoint; index <= 10000; index += input {
		fmt.Println(index)
	}
}
