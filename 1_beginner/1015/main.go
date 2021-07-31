package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func distance(point1 Point, point2 Point) float64 {
	var parcialResult, parcialResult1, parcialResult2 float64

	parcialResult1 = math.Pow(point1.x, 2) - 2*(point1.x*point2.x) + math.Pow(point2.x, 2)
	parcialResult2 = math.Pow(point1.y, 2) - 2*(point1.y*point2.y) + math.Pow(point2.y, 2)

	parcialResult = parcialResult1 + parcialResult2

	return math.Sqrt(parcialResult)
}

func main() {
	var result float64
	var point1, point2 Point

	fmt.Scanf("%f %f", &point1.x, &point1.y)
	fmt.Scanf("%f %f", &point2.x, &point2.y)

	result = distance(point1, point2)

	fmt.Printf("%.4f\n", result)
}
