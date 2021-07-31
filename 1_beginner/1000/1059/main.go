package main

import (
	"fmt"
)

func main() {
	for index := 1; index <= 100; index++ {
		if index%2 != 0 {
			continue
		}

		fmt.Println(index)
	}
}
