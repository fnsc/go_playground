package main

import "fmt"

func main() {
	input := []int{0, 0}
	newLine := 1

	fmt.Scanf("%d %d", &input[0], &input[1])

	for index := 1; index <= input[1]; index++ {
		fmt.Print(index)

		if newLine == input[0] || index == input[1] {
			fmt.Printf("\n")

			newLine = 1

			continue
		}

		fmt.Print(" ")
		newLine++
	}
}
