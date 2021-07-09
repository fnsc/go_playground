package main

import (
	"fmt"
)

type Result struct {
	year, month, day int
}

/*
 * I'm not using math.Round because it was introduced
 * on version 1.10 and, the code must run on version 1.8
 */
func round(unit float64) float64 {
	intPart := int(unit)
	decimalPart := unit - float64(intPart)

	if decimalPart >= 0.1 {
		return float64(intPart) + 1
	}

	return float64(intPart)
}

func convert(reference float64, factor float64, input int) int {
	unit := float64(input) / float64(reference)
	intPart := int(unit)
	unit -= float64(intPart)
	unit *= float64(factor)

	if reference == factor {
		return int(round(unit))
	}

	return int(unit)
}

func calculate(days int) Result {
	var result Result
	const MONTH = 30.416666667
	const YEAR = 12.166666667

	if days < 30 {
		result.year = 0
		result.month = 0
		result.day = days

		return result
	}

	if days == 30 {
		result.year = 0
		result.month = 1
		result.day = 0

		return result
	}

	if days == 365 {
		result.year = 1
		result.month = 0
		result.day = 0

		return result
	}

	if days >= 30 && days < 365 {

		result.year = 0
		result.month = int(float64(days) / MONTH)
		result.day = convert(MONTH, MONTH, days)

		return result
	}

	result.year = days / 365
	result.month = convert(365, YEAR, days)
	result.day = convert(MONTH, MONTH, days)

	return result
}

func main() {
	var input int

	fmt.Scanf("%d", &input)

	result := calculate(input)

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", result.year, result.month, result.day)
}
