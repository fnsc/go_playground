package main

import "fmt"

func main() {
	input := 0
	entries := []int{}

	for input >= 0 {
		read(&input, &entries)
	}

	fmt.Printf("%.2f\n", calculateAverage(&entries))
}

func read(input *int, entries *[]int) {
	fmt.Scanf("%d", input)

	addNewEntry(entries, input)
}

func addNewEntry(entries *[]int, input *int) {
	if *input < 0 {
		return
	}

	*entries = append(*entries, *input)
}

func calculateAverage(entries *[]int) float64 {
	agesSum := 0
	entriesQuantity := len(*entries)

	for _, age := range *entries {
		agesSum += age
	}

	return float64(agesSum) / float64(entriesQuantity)
}
