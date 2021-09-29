package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{0, 0}

	fmt.Scanf("%d", &input[0])
	fmt.Scanf("%d", &input[1])

	sort.Ints(input[:])

	fmt.Println(calculate(input))
}

func calculate(input []int) int {
	total := 0

	for index := input[0]; index <= input[1]; index++ {
		if index%13 != 0 {
			total += index
		}
	}

	return total
}
