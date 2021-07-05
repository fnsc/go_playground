package main

import (
	"fmt"
	"math"
)

func greatest(input1 int, input2 int) int {
	var result float64 = (float64(input1) + float64(input2) + math.Abs(float64(input1)-float64(input2))) / 2

	return int(result)
}

func main() {
	var result, input1, input2, input3 int

	fmt.Scanf("%d %d %d", &input1, &input2, &input3)

	result = greatest(input1, input2)
	result = greatest(input3, result)

	fmt.Printf("%d eh o maior\n", result)
}
