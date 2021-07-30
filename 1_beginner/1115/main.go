package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	input := []int{0, 0}

	for 0 < 1 {
		fmt.Scanf("%d %d", &input[0], &input[1])

		if canFinish(input) {
			return
		}

		getPosition(input)
	}
}

func canFinish(input []int) bool {
	return input[0] == 0 || input[1] == 0
}

func getPosition(input []int) {
	if input[0] > 0 && input[1] > 0 {
		printPosition("primeiro")

		return
	}

	if input[0] < 0 && input[1] > 0 {
		printPosition("segundo")

		return
	}

	if input[0] < 0 && input[1] < 0 {
		printPosition("terceiro")

		return
	}

	printPosition("quarto")
}

func printPosition(message string) {
	fmt.Println(message)
}
