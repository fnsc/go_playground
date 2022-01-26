package main

import "fmt"

func main() {
	input := 0
	number := 0

	read(&input)

	for i := 0; i < input; i++ {
		read(&number)

		printResult(&number, isPrimeNumber(&number))
	}
}

func read(input *int) {
	fmt.Scanf("%d", input)
}

func isPrimeNumber(number *int) bool {
	if *number == 2 {
		return true
	}

	if *number%2 == 0 {
		return false
	}

	divisors := []int{}
	i := 1

	for i <= *number {
		if *number%i != 0 {
			i += 2

			continue
		}

		divisors = append(divisors, i)
		i += 2
	}

	return len(divisors) == 2
}

func printResult(number *int, result bool) {
	if result {
		fmt.Println(*number, "eh primo")

		return
	}

	fmt.Println(*number, "nao eh primo")
}
