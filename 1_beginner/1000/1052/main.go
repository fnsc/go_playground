package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanf("%d", &input)

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	for index, month := range months {
		if !(input == index+1) {
			continue
		}

		fmt.Println(month)

		return
	}
}
