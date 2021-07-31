package main

import (
	"fmt"
)

type Product struct {
	code     int
	quantity int
	value    float64
}

func calculate(product1 Product, product2 Product) float64 {
	return (product1.value * float64(product1.quantity)) + (product2.value * float64(product2.quantity))
}

func main() {
	var result float64
	var product1, product2 Product

	fmt.Scanf("%d %d %f", &product1.code, &product1.quantity, &product1.value)
	fmt.Scanf("%d %d %f", &product2.code, &product2.quantity, &product2.value)

	result = calculate(product1, product2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", result)
}
