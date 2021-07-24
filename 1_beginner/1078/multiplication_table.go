package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanf("%d", &input)

	fmt.Printf("1 x %d = %d\n2 x %d = %d\n3 x %d = %d\n4 x %d = %d\n5 x %d = %d\n6 x %d = %d\n7 x %d = %d\n8 x %d = %d\n9 x %d = %d\n10 x %d = %d\n",
		input, input*1, input, input*2, input, input*3, input, input*4, input, input*5, input, input*6, input, input*7, input, input*8, input, input*9, input, input*10)
}
