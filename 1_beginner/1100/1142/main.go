package main

import "fmt"

func main() {
	input := 0
	cursor := 0

	fmt.Scanf("%d", &input)

	for index := 1; index <= input; index++ {
		cursor++
		first := cursor
		cursor++
		second := cursor
		cursor++
		third := cursor
		cursor++
		fmt.Printf("%d %d %d PUM\n", first, second, third)
	}
}
