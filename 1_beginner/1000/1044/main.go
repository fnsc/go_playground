package main

import (
	"fmt"
	"sort"
)

func areMultiples(input [2]int) bool {
	return input[1]%input[0] == 0
}

func main() {
	var input [2]int

	fmt.Scanf("%d %d", &input[0], &input[1])

	sort.Ints(input[:])

	if areMultiples(input) {
		fmt.Println("Sao Multiplos")

		return
	}

	fmt.Println("Nao sao Multiplos")
}
