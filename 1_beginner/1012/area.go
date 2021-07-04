package main

import (
	"fmt"
	"math"
)

func triangleArea(base float64, height float64) float64 {
	return base * height / 2
}

func circleArea(radius float64) float64 {
	const PI = 3.14159

	return PI * math.Pow(radius, 2)
}

func trapeziumArea(base1 float64, base2 float64, height float64) float64 {
	return ((base1 + base2) / 2) * height
}

func squareArea(side float64) float64 {
	return side * side
}

func rectangleArea(base float64, height float64) float64 {
	return base * height
}

func main() {
	var input1, input2, input3,
		triangle, circle, trapezium,
		square, rectangle float64

	fmt.Scanf("%f %f %f", &input1, &input2, &input3)

	triangle = triangleArea(input1, input3)
	circle = circleArea(input3)
	trapezium = trapeziumArea(input1, input2, input3)
	square = squareArea(input2)
	rectangle = rectangleArea(input1, input2)

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", triangle, circle, trapezium, square, rectangle)
}
