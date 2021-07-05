package main

import (
	"fmt"
)

func main() {
	var result, fuel float64
	var distance int

	fmt.Scanf("%d", &distance)
	fmt.Scanf("%f", &fuel)

	result = float64(distance) / fuel

	fmt.Printf("%.3f km/l\n", result)
}
