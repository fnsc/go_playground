package main

import (
	"fmt"
)

func main() {
	var quantity int

	fmt.Scanf("%d", &quantity)

	numbers := make([]int, quantity)

	for index := 0; index < quantity; index++ {
		fmt.Scanf("%d", &numbers[index])
	}

	calculate(numbers)
}

func calculate(numbers []int) {
	for _, number := range numbers {
		printResult(number)
	}
}

func printResult(number int) {
	if number == 0 {
		fmt.Println("NULL")

		return
	}

	if number%2 == 0 {
		if number > 0 {
			fmt.Println("EVEN POSITIVE")
			return
		}

		fmt.Println("EVEN NEGATIVE")

		return
	}

	if number > 0 {
		fmt.Println("ODD POSITIVE")

		return
	}

	fmt.Println("ODD NEGATIVE")
}
