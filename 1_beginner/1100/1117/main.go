package main

import "fmt"

func main() {
	grades := readInput()
	average := calculate(grades)
	printResult(average)
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
