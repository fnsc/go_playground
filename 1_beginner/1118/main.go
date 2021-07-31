package main

import "fmt"

func main() {
	for 0 < 1 {
		grades := readInput()
		average := calculate(grades)
		printResult(average)

		action := askForAction()

		if !startOver(action) {
			return
		}
	}
}

func readInput() []float64 {
	counter := 2
	input := 0.0
	grades := make([]float64, 0)

	for counter > 0 {
		fmt.Scanf("%f", &input)

		if !isValid(input) {
			fmt.Println("nota invalida")

			continue
		}

		grades = append(grades, input)
		counter--
	}

	return grades
}

func isValid(grade float64) bool {
	return grade >= 0 && grade <= 10
}

func calculate(grades []float64) float64 {
	return (grades[0] + grades[1]) / 2
}

func printResult(average float64) {
	fmt.Printf("media = %.2f\n", average)
}

func askForAction() int {
	input := 0

	fmt.Println("novo calculo (1-sim 2-nao)")

	fmt.Scanf("%d", &input)

	return input
}

func startOver(action int) bool {
	if action == 1 {
		return true
	}

	if action == 2 {
		return false
	}

	action = askForAction()

	return startOver(action)
}
