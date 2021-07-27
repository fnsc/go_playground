package main

import (
	"fmt"
)

func main() {
	j := 5
	for i := 1; i <= 9; i += 2 {
		fmt.Printf("I=%d J=%d\nI=%d J=%d\nI=%d J=%d\n", i, j+2, i, j+1, i, j)

		j += 2
	}
}
