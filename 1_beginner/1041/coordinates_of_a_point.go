package main

import "fmt"

type Point struct {
	x, y float64
}

func belongsToQ1(point Point) bool {
	return point.x >= 0.000000001 && point.y >= 0.000000001
}

func belongsToQ2(point Point) bool {
	return point.x <= -0.00000001 && point.y >= 0.000000001
}

func belongsToQ3(point Point) bool {
	return point.x <= -0.000000001 && point.y <= -0.000000001
}

func belongsToQ4(point Point) bool {
	return point.x >= 0.000000001 && point.y <= -0.000000001
}

func xAxis(point Point) bool {
	return point.x != 0.0 && point.y == 0.0
}

func yAxis(point Point) bool {
	return point.x == 0.0 && point.y != 0.0
}

func main() {
	var point Point

	fmt.Scanf("%f %f", &point.x, &point.y)

	if belongsToQ1(point) {
		fmt.Println("Q1")

		return
	}

	if belongsToQ2(point) {
		fmt.Println("Q2")

		return
	}

	if belongsToQ3(point) {
		fmt.Println("Q3")

		return
	}

	if belongsToQ4(point) {
		fmt.Println("Q4")

		return
	}

	if xAxis(point) {
		fmt.Println("Eixo X")

		return
	}

	if yAxis(point) {
		fmt.Println("Eixo Y")

		return
	}

	fmt.Println("Origem")
}
