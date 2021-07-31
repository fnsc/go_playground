package main

import (
	"fmt"
)

func main() {
	const AVERAGE = 12

	var result float64
	var hours, speed, distance int

	fmt.Scanf("%d", &hours)
	fmt.Scanf("%d", &speed)

	distance = hours * speed
	result = float64(distance) / AVERAGE

	fmt.Printf("%.3f\n", result)
}
