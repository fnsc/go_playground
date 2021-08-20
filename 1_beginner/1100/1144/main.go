package main

import (
	"fmt"
)

func main() {
	input := 0

	fmt.Scanf("%d", &input)

	for index := 1; index <= input; index++ {
		firstTerm := index * index
		secondTerm := firstTerm * index

		fmt.Printf("%d %d %d\n", index, firstTerm, secondTerm)
		fmt.Printf("%d %d %d\n", index, firstTerm+1, secondTerm+1)
	}
}
