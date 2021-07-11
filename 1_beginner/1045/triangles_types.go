package main

import (
	"fmt"
	"math"
	"sort"
)

type Triangle struct {
	sides [3]float64
}

func isInvalid(triangle Triangle) bool {
	return triangle.sides[0] >= triangle.sides[1]+triangle.sides[2]
}

func isRight(triangle Triangle) bool {
	return math.Pow(triangle.sides[0], 2) == math.Pow(triangle.sides[1], 2)+math.Pow(triangle.sides[2], 2)
}

func isObtuse(triangle Triangle) bool {
	return math.Pow(triangle.sides[0], 2) > math.Pow(triangle.sides[1], 2)+math.Pow(triangle.sides[2], 2)
}

func isAcute(triangle Triangle) bool {
	return math.Pow(triangle.sides[0], 2) < math.Pow(triangle.sides[1], 2)+math.Pow(triangle.sides[2], 2)
}

func isEquilateral(triangle Triangle) bool {
	return triangle.sides[0] == triangle.sides[1] && triangle.sides[0] == triangle.sides[2]
}

func isIsosceles(triangle Triangle) bool {
	return (triangle.sides[0] == triangle.sides[1] && triangle.sides[0] != triangle.sides[2]) ||
		(triangle.sides[1] == triangle.sides[2] && triangle.sides[1] != triangle.sides[0])
}

func printResult(message string) {
	fmt.Println(message)
}

func main() {
	var triangle Triangle

	fmt.Scanf("%f %f %f", &triangle.sides[0], &triangle.sides[1], &triangle.sides[2])

	sort.Sort(sort.Reverse(sort.Float64Slice(triangle.sides[:])))

	if isInvalid(triangle) {
		printResult("NAO FORMA TRIANGULO")

		return
	}

	if isRight(triangle) {
		printResult("TRIANGULO RETANGULO")
	}

	if isObtuse(triangle) {
		printResult("TRIANGULO OBTUSANGULO")
	}

	if isAcute(triangle) {
		printResult("TRIANGULO ACUTANGULO")
	}

	if isEquilateral(triangle) {
		printResult("TRIANGULO EQUILATERO")
	}

	if isIsosceles(triangle) {
		printResult("TRIANGULO ISOSCELES")
	}
}
