package main

import (
	"fmt"
)

type Time struct {
	hour, minute, second int
}

func convert(input int) Time {
	// 1 min has 60 seconds
	// 1 hour has 3600 seconds
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

		seconds := float64(input) / 60.0
		intPart := int(seconds)
		seconds -= float64(intPart)

		seconds *= 60

		time.second = int(seconds)

		return time
	}

	time.hour = input / 3600

	minutes := float64(input) / 3600.0
	intPart := int(minutes)
	minutes -= float64(intPart)

	time.minute = int(minutes)

	seconds := float64(input) / 60.0
	intPart = int(seconds)
	seconds -= float64(intPart)

	seconds *= 60

	time.second = int(seconds)

	return time
}

func main() {
	var input int

	fmt.Scanf("%d", &input)

	result := convert(input)

	fmt.Printf("%d:%d:%d\n", result.hour, result.minute, result.second)
}
