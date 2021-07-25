package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i += 2 {
		fmt.Printf("I=%d J=7\nI=%d J=6\nI=%d J=5\n", i, i, i)
	}
}
