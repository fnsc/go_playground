package main

import "fmt"

func main() {
	var quantity int

	fmt.Scanf("%d", &quantity)

	numbers := make([]int, quantity)

	for index := 0; index < quantity; index++ {
		fmt.Scanf("%d", &numbers[index])
	}

	in, out := calculate(numbers)

	fmt.Printf("%d in\n%d out\n", in, out)
}

func calculate(numbers []int) (int, int) {
	in, out := 0, 0

	for _, number := range numbers {
		if number >= 10 && number <= 20 {
			in++
			continue
		}

		out++
	}

	return in, out
}
