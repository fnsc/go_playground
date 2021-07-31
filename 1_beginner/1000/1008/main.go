package main

import (
	"fmt"
)

func calculate(hours int, salary float64) float64 {
	return float64(hours) * salary
}

func main() {
	var number, hours int
	var price, salary float64

	fmt.Scanf("%d", &number)
	fmt.Scanf("%d", &hours)
	fmt.Scanf("%f", &price)

	salary = calculate(hours, price)

	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", number, salary)
}
