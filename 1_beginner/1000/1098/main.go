package main

import (
	"fmt"
)

func main() {
	for i := 0.0; i <= 20; i += 2 {
		reference := i / 10
		integerPart := int(reference)

		if reference-float64(integerPart) == 0.0 {
			fmt.Printf("I=%d J=%d\nI=%d J=%d\nI=%d J=%d\n", integerPart, 1+integerPart, integerPart, 2+integerPart, integerPart, 3+integerPart)

			continue
		}

		fmt.Printf("I=%.1f J=%.1f\nI=%.1f J=%.1f\nI=%.1f J=%.1f\n", reference, 1+reference, reference, 2+reference, reference, 3+reference)
	}
}
