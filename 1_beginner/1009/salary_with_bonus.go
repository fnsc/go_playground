package main

import (
	"fmt"
)

func calculate(salary float64, sale float64) float64 {
	const BONUS = 0.15

	return salary + (sale * BONUS)
}

func main() {
	var name string
	var sale, salary float64

	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &salary)
	fmt.Scanf("%f", &sale)

	salary = calculate(salary, sale)

	fmt.Printf("TOTAL = R$ %.2f\n", salary)
}
