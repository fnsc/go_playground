package main

import "fmt"

func main() {
	input := [10]int{}
	fmt.Scanf("%d", &input[0])

	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Printf("N[%d] = %d\n", i, input[i])

			continue
		}

		input[i] = input[i-1] * 2

		fmt.Printf("N[%d] = %d\n", i, input[i])
	}
}
