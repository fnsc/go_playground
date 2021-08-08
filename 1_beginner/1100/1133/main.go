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

	printResult(input)
}

func printResult(input []int) {
	for index := input[0] + 1; index < input[1]; index++ {
		rest := index % 5
		if rest == 2 || rest == 3 {
			fmt.Println(index)
		}
	}
}
