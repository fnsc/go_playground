package main

import "fmt"

type Triangle struct {
	side1, side2, side3 float64
}

func isValid(triangle Triangle) bool {
	return triangle.side1+triangle.side2 > triangle.side3 &&
		triangle.side1+triangle.side3 > triangle.side2 &&
		triangle.side2+triangle.side3 > triangle.side1
}

func getPerimeter(triangle Triangle) float64 {
	return triangle.side1 + triangle.side2 + triangle.side3
}

func trapeziumArea(triangle Triangle) float64 {
	return ((triangle.side1 + triangle.side2) / 2) * triangle.side3
}

func main() {
	var triangle Triangle

	fmt.Scanf("%f %f %f", &triangle.side1, &triangle.side2, &triangle.side3)

	if isValid(triangle) {
		result := getPerimeter(triangle)
		fmt.Printf("Perimetro = %.1f\n", result)

		return
	}

	result := trapeziumArea(triangle)
	fmt.Printf("Area = %.1f\n", result)
}
