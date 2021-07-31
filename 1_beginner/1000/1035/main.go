package main

import (
	"fmt"
)

type Input struct {
	a, b, c, d int
}

func greater(input Input) bool {
	return input.b > input.c && input.d > input.a
}

func sumGreater(input Input) bool {
	return input.c+input.d > input.a+input.b
}

func isPositive(input Input) bool {
	return input.c >= 0 && input.d >= 0
}

func isEven(input Input) bool {
	return input.a%2 == 0
}

func main() {
	var input Input

	fmt.Scanf("%d %d %d %d", &input.a, &input.b, &input.c, &input.d)

	if greater(input) && sumGreater(input) && isPositive(input) && isEven(input) {
		fmt.Printf("Valores aceitos\n")

		return
	}

	fmt.Printf("Valores nao aceitos\n")

}
