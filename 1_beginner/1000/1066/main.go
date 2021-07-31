package main

import "fmt"

func main() {
	var input [5]int

	fmt.Scanf("%d", &input[0])
	fmt.Scanf("%d", &input[1])
	fmt.Scanf("%d", &input[2])
	fmt.Scanf("%d", &input[3])
	fmt.Scanf("%d", &input[4])

	calculate(input)
}

func calculate(input [5]int) {
	evenCounter, oddCounter, positiveCounter, negativeCouter := 0, 0, 0, 0

	for _, number := range input {
		if number%2 == 0 {
			evenCounter++
			if number > 0 {
				positiveCounter++
				continue
			}
			if number < 0 {
				negativeCouter++
			}
			continue
		}

		oddCounter++
		if number > 0 {
			positiveCounter++
			continue
		}
		if number < 0 {
			negativeCouter++
		}
	}

	printResult(evenCounter, oddCounter, positiveCounter, negativeCouter)
}

func printResult(evenCounter, oddCounter, positiveCounter, negativeCouter int) {
	fmt.Printf("%d valor(es) par(es)\n%d valor(es) impar(es)\n%d valor(es) positivo(s)\n%d valor(es) negativo(s)\n", evenCounter, oddCounter, positiveCounter, negativeCouter)
}
