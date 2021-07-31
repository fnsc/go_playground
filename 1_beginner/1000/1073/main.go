package main

import (
	"fmt"
	"math"
)

func main() {
	var input int

	fmt.Scanf("%d", &input)

	result := calculate(input)

	printResult(result)
}

type Number struct {
	operator int
	pow      float64
}

func calculate(input int) []Number {
	var number Number
	numbers := make([]Number, 0)

	for index := 2; index <= input; index++ {
		if index%2 != 0 {
			continue
		}

		number.operator = index
		number.pow = math.Pow(float64(index), 2)

		numbers = append(numbers, number)
	}

	return numbers
}

func printResult(result []Number) {
	for _, number := range result {
		fmt.Printf("%d^2 = %d\n", number.operator, int(number.pow))
	}
}
