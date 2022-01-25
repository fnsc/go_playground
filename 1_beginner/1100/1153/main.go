package main

import "fmt"

func main() {
	input := 0

	read(&input)

	fmt.Println(factorial(input))
}

func read(input *int) {
	fmt.Scanf("%d", input)

	if *input <= 0 || *input >= 13 {
		read(input)
	}
}

func factorial(input int) int {
	if input <= 1 {
		return 1
	}

	return input * factorial(input-1)
}
