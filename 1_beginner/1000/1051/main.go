package main

import (
	"fmt"
)

func calculate(salary float64) float64 {
	var totalTaxes float64 = 0

	if salary >= 4500.01 {
		totalTaxes += 1000 * 0.08
		totalTaxes += 1500 * 0.18
		totalTaxes += (salary - 4500) * 0.28

		return totalTaxes
	}

	if salary >= 3000.01 && salary <= 4500 {
		totalTaxes += 1000 * 0.08
		totalTaxes += (salary - 3000) * 0.18

		return totalTaxes
	}

	if salary >= 2000.01 && salary <= 3000 {
		totalTaxes += (salary - 2000) * 0.08

		return totalTaxes
	}

	return 0
}

func printResult(taxes float64) {
	if taxes == 0 {
		fmt.Println("Isento")
		return
	}

	fmt.Printf("R$ %.2f\n", taxes)
}

func main() {
	var salary float64

	fmt.Scanf("%f", &salary)

	result := calculate(salary)

	printResult(result)
}
