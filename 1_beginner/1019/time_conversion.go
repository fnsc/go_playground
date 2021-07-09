package main

import (
	"fmt"
)

type Time struct {
	hour, minute, second int
}

/* I'm not using math.Round because it was introduced
 * on version 1.10 and, the code must run on version 1.8
 */
func round(unit float64) float64 {
	intPart := int(unit)
	decimalPart := unit - float64(intPart)

	if decimalPart >= 0.5 {
		return float64(intPart) + 1
	}

	return float64(intPart)
}

func convert(reference int, input int) int {
	unit := float64(input) / float64(reference)
	intPart := int(unit)
	unit -= float64(intPart)
	unit *= 60

	if reference == 60 {
		return int(round(unit))
	}

	return int(unit)
}

func calculate(input int) Time {
	var time Time

	if input < 60 {
		time.hour = 0
		time.minute = 0
		time.second = input

		return time
	}

	if input >= 60 && input < 3600 {
		time.hour = 0
		time.minute = input / 60
		time.second = convert(60, input)

		return time
	}

	time.hour = input / 3600
	time.minute = convert(3600, input)
	time.second = convert(60, input)

	return time
}

func main() {
	var input int

	fmt.Scanf("%d", &input)

	result := calculate(input)

	fmt.Printf("%d:%d:%d\n", result.hour, result.minute, result.second)
}
