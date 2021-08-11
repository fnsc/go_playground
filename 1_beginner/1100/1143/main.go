package main

import (
	"fmt"
	"math"
)

func main() {
	input := 0
	cursor := 1.0

	fmt.Scanf("%d", &input)

	for index := 1; index <= input; index++ {
		square := math.Pow(cursor, 2)
		cubic := math.Pow(cursor, 3)

		fmt.Printf("%.0f %.0f %.0f\n", cursor, square, cubic)

		cursor++
	}
}
