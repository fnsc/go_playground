package main

import "testing"
import "github.com/stretchr/testify/assert"
import "io/ioutil"
import "os"
import "log"

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
		assert.Equal(test, scenario.expected, scenario.input, "Expecting 1")
	}
}

func TestShouldReturnOneWhenNilIsGiven(test *testing.T) {
	// Set
	var input int
	expected := 1

	// Action
	isLowerThanOne(&input)

	// Assertions
	assert.Equal(test, expected, input, "Expecting 1")
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

func Benchmark_main(test *testing.B) {

	for i := 0; i <= test.N; i++ {
		content := []byte("-1\n2\n-3\n4\n-5\n6\n-7\n8\n-9\n\n")
		tmpfile, err := ioutil.TempFile("", "example")
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write(content); err != nil {
			log.Fatal(err)
		}

		if _, err := tmpfile.Seek(0, 0); err != nil {
			log.Fatal(err)
		}

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = tmpfile
		main()

		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	}

}

func TestShouldRunTheMainFunc(test *testing.T) {
	content := []byte("-1\n2\n-3\n4\n-5\n6\n-7\n8\n-9\n\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = tmpfile
	main()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
