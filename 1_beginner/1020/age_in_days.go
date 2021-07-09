package main

import (
	"fmt"
)

type Result struct {
	year, month, day int
}

func calculate(days int) Result {
	var result Result
	result.year = days / 365
	days %= 365
	result.month = days / 30
	days %= 30
	result.day = days

	return result
}

func main() {
	var input int

	fmt.Scanf("%d", &input)

	result := calculate(input)

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", result.year, result.month, result.day)
}
