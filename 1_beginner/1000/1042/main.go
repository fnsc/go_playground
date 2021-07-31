package main

import (
	"fmt"
	"sort"
)

type Number struct {
	inputs  [3]int
	ordered [3]int
}

func main() {
	var numbers Number

	fmt.Scanf("%d %d %d", &numbers.inputs[0], &numbers.inputs[1], &numbers.inputs[2])

	numbers.ordered = numbers.inputs

	sort.Ints(numbers.ordered[:])

	fmt.Printf(
		"%d\n%d\n%d\n\n%d\n%d\n%d\n",
		numbers.ordered[0], numbers.ordered[1], numbers.ordered[2], numbers.inputs[0], numbers.inputs[1], numbers.inputs[2])
}
