package main

import (
	"fmt"
	"strings"
)

func main() {
	input := 0

	read(&input)

	fibonacci := getFibonacci(&input)

	printFibonacci(&fibonacci)
}

func read(input *int) {
	fmt.Scanf("%d", input)

	if *input <= 0 || *input >= 46 {
		read(input)
	}
}

func getFibonacci(target *int) []int {
	result := []int{0, 1}
	nextValue := 0

	for len(result) < *target {
		nextValue = result[len(result)-2] + result[len(result)-1]

		result = append(result, nextValue)
	}

	return result
}

func printFibonacci(fibonacci *[]int) {
	result := ""

	for _, number := range *fibonacci {
		result += fmt.Sprintf("%d ", number)
	}

	result = strings.TrimRight(result, " ")

	fmt.Println(result)
}
