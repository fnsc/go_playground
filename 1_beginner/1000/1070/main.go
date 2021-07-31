package main

import "fmt"

func main() {
	var input int

	fmt.Scanf("%d", &input)

	counter := 0
	for index := input; index <= input+11; index++ {
		if index%2 == 0 {
			continue
		}

		if counter < 6 {
			fmt.Println(index)
		}

		counter++
	}

}
