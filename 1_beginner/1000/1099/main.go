package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int

	fmt.Scanf("%d", &input)

	numbers := make([][3]int, input)

	for index, couple := range numbers {
		fmt.Scanf("%d %d", &couple[0], &couple[1])

		sort.Ints(couple[:])

		if couple[2]-couple[1] <= 1 {
			continue
		}

		calculate(&couple)

		numbers[index] = couple
	}

	printResult(numbers)
}

func calculate(numbers *[3]int) {
	for index := numbers[1] + 1; index < numbers[2]; index++ {
		if index%2 == 0 {
			continue
		}

		numbers[0] += index
	}
}

func printResult(numbers [][3]int) {
	for _, number := range numbers {
		fmt.Println(number[0])
	}
}
