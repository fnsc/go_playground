package main

import "fmt"

func main() {
	input := [20]int{}
	aux1 := 19
	aux2 := 9

	for i := 0; i < 20; i++ {
		if i <= 9 {
			fmt.Scanf("%d", &input[aux1])
			aux1--

			continue
		}

		fmt.Scanf("%d", &input[aux2])
		aux2--
	}

	for index, value := range input {
		fmt.Printf("N[%d] = %d\n", index, value)
	}
}
