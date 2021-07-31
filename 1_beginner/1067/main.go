package main

import "fmt"

func main() {
	var input int

	fmt.Scanf("%d", &input)

	for index := 1; index <= input; index++ {
		if index%2 == 0 {
			continue
		}

		fmt.Println(index)
	}
}
