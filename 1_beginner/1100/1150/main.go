package main

import "fmt"

func main() {
	startingPoint := 0
	target := 0

	read(&startingPoint)

	for target <= startingPoint {
		read(&target)
	}

	fmt.Println(getResult(&startingPoint, &target))
}

func read(input *int) {
	fmt.Scanf("%d", input)
}

func getResult(startingPoint, target *int) int {
	result := 1

	for *startingPoint <= *target {
		*startingPoint++
		*startingPoint += *startingPoint

		result++
	}

	result++

	return result
}
