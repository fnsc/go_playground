package main

import "fmt"

func main() {
	input := 0

	fmt.Scanf("%d", &input)

	if input%2 == 0 {
		result := getEvenDivisors(&input)
		printResult(result)

		return
	}

	result := getOddDivisors(&input)

	printResult(result)
}

func getEvenDivisors(input *int) *[]int {
	result := []int{}

	for i := 1; i <= *input; i++ {
		if *input%i != 0 {
			continue
		}

		result = append(result, i)
	}

	return &result
}

func getOddDivisors(input *int) *[]int {
	result := []int{}
	i := 1

	for i <= *input {
		if *input%i != 0 {
			i += 2

			continue
		}

		result = append(result, i)

		i += 2
	}

	return &result
}

func printResult(result *[]int) {
	for _, divisor := range *result {
		fmt.Println(divisor)
	}
}
