package main

import (
	"fmt"
	"sort"
)

func main() {
	var input [2]int

	fmt.Scanf("%d", &input[0])
	fmt.Scanf("%d", &input[1])

	sort.Ints(input[:])

	result := calculate(input)

	fmt.Println(result)
}

func calculate(input [2]int) int {
	sum := 0
	for index := input[0] + 1; index < input[1]; index++ {
		if index%2 == 0 {
			continue
		}

		sum += index
	}

	return sum
}
