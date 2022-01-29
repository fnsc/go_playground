package main

import "fmt"

func main() {
	input := [10]int{}

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &input[i])

		isLowerThanOne(&input[i])
	}

	printResult(&input)
}

func isLowerThanOne(input *int) {
	if *input <= 0 {
		*input = 1
	}
}

func printResult(input *[10]int) {
	for index, value := range *input {
		fmt.Printf("X[%d] = %d\n", index, value)
	}
}
