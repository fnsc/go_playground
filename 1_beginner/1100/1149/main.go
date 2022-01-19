package main

import "fmt"

func main() {
	total := readInput()

	fmt.Println(total)
}

func calculateSum(userEntries []int) int {
	total := 0
	for i := 0; i < userEntries[len(userEntries)-1]; i++ {
		total += userEntries[0] + i
	}

	return total
}

func readInput() int {
	input := 0
	userEntries := []int{}

	for 1 == 1 {
		fmt.Scanf("%d", &input)
		userEntries = append(userEntries, input)

		if len(userEntries) >= 2 && input > 0 {
			return calculateSum(userEntries)
		}
	}

	return 0
}
