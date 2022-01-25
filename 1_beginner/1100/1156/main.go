package main

import "fmt"

func main() {
	result := 0.0
	i := 1.0
	dividedBy := 1.0

	for i <= 39.0 {
		result += i / dividedBy

		i += 2
		dividedBy *= 2
	}

	fmt.Printf("%.2f\n", result)
}
