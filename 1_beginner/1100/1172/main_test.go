package main

import "testing"

type lowerThanOneDataProvider struct {
	input    int
	expected int
}

func TestShouldValidateIfTheUserInputIsGreaterThanZero(test *testing.T) {
	// Set
	scenarios := []lowerThanOneDataProvider{
		lowerThanOneDataProvider{input: 0, expected: 1},
		lowerThanOneDataProvider{input: -1, expected: 1},
		lowerThanOneDataProvider{expected: 1},
	}

	// Action
	for _, scenario := range scenarios {
		isLowerThanOne(&scenario.input)

		// Assertions
		if scenario.input != 1 {
			test.Error("Expected:", scenario.expected, "Got:", scenario.input)
		}
	}
}

func Example_printResult() {
	input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printResult(&input)
	// Output:
	// X[0] = 1
	// X[1] = 2
	// X[2] = 3
	// X[3] = 4
	// X[4] = 5
	// X[5] = 6
	// X[6] = 7
	// X[7] = 8
	// X[8] = 9
	// X[9] = 10
}

func Example_main() {
	// input := [10]int{-1, 2, -3, 4, -5, 6, -7, 8, -9, 10}
	main()
	// Output:
	// X[0] = 1
	// X[1] = 2
	// X[2] = 1
	// X[3] = 4
	// X[4] = 1
	// X[5] = 6
	// X[6] = 1
	// X[7] = 8
	// X[8] = 1
	// X[9] = 10
}
