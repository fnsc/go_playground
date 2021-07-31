package main

import "fmt"

func isValid(input float64) bool {
	return input >= 0 && input <= 100
}

func isBetweenZeroToTwentyFive(input float64) bool {
	return input >= 0 && input <= 25.0000
}

func isBetweenTwentyFiveToFifty(input float64) bool {
	return input >= 25.00001 && input <= 50.0000000
}

func isBetweenFiftyToSeventyFive(input float64) bool {
	return input >= 50.00000001 && input <= 75.0000000
}

func isBetweenSeventyFiveToHundred(input float64) bool {
	return input >= 75.00000001 && input <= 100.0000000
}

func main() {
	var input float64

	fmt.Scanf("%f", &input)

	if !isValid(input) {
		fmt.Println("Fora de intervalo")

		return
	}

	if isBetweenZeroToTwentyFive(input) {
		fmt.Println("Intervalo [0,25]")

		return
	}

	if isBetweenTwentyFiveToFifty(input) {
		fmt.Println("Intervalo (25,50]")

		return
	}

	if isBetweenFiftyToSeventyFive(input) {
		fmt.Println("Intervalo (50,75]")

		return
	}

	if isBetweenSeventyFiveToHundred(input) {
		fmt.Println("Intervalo (75,100]")

		return
	}
}
