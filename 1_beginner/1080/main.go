package main

import (
	"fmt"
)

func main() {
	var input int
	result := [2]int{0, 0}

	for index := 1; index <= 100; index++ {
		fmt.Scanf("%d", &input)

		if input > result[0] {
			result[0] = input
			result[1] = index
		}
	}

	fmt.Printf("%d\n%d\n", result[0], result[1])
}
