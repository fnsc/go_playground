package main

import "fmt"

func calculate(start, end int) int {
	const DAY_DURATION = 24
	if start == end {
		return DAY_DURATION
	}

	if start < end {
		return end - start
	}

	return DAY_DURATION - start + end
}

func main() {
	var start, end int

	fmt.Scanf("%d %d", &start, &end)

	result := calculate(start, end)

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", result)
}
