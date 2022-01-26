package main

import "fmt"

func main() {
	input := 0
	number := 0

	read(&input, 20)

	for i := 0; i < input; i++ {
		read(&number, 1000000000)

		result := calculate(&number)
		printResult(&number, result)
	}
}

func read(input *int, limit int) {
	fmt.Scanf("%d", input)

	if *input < 1 || *input > limit {
		read(input, limit)
	}
}

func calculate(number *int) *int {
	divisorsSum := 0

	for i := 1; i <= *number; i++ {
		if *number%i != 0 || *number == i {
			continue
		}

		divisorsSum += i
	}

	return &divisorsSum
}

func printResult(number, result *int) {
	if *number == *result {
		fmt.Println(*number, "eh perfeito")

		return
	}

	fmt.Println(*number, "nao eh perfeito")
}
