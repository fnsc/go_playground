package main

import "fmt"

func main() {
	input := 0.0
	i := 0

	for i < 100 {
		fmt.Scanf("%f", &input)

		if input <= 10 {
			fmt.Printf("A[%d] = %.1f\n", i, input)
		}

		i++
	}
}
