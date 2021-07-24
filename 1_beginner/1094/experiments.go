package main

import (
	"fmt"
)

func main() {
	var input, quantity, total int
	var kind string
	var rabit, rat, frog Animal

	fmt.Scanf("%d", &input)

	for index := 1; index <= input; index++ {
		fmt.Scanf("%d %s", &quantity, &kind)

		total += quantity

		if kind == "C" {
			rabit.quantity += quantity

			continue
		}

		if kind == "R" {
			rat.quantity += quantity

			continue
		}

		frog.quantity += quantity
	}

	rabit.percentage = calculatePercentage(rabit, total)
	rat.percentage = calculatePercentage(rat, total)
	frog.percentage = calculatePercentage(frog, total)

	printResult(total, rabit, rat, frog)
}

type Animal struct {
	quantity   int
	percentage float64
}

func calculatePercentage(animal Animal, total int) float64 {
	animal.percentage = float64(animal.quantity) * 100 / float64(total)

	return animal.percentage
}

func printResult(total int, rabit, rat, frog Animal) {
	fmt.Printf("Total: %d cobaias\nTotal de coelhos: %d\nTotal de ratos: %d\nTotal de sapos: %d\nPercentual de coelhos: %.2f %%\nPercentual de ratos: %.2f %%\nPercentual de sapos: %.2f %%\n",
		total, rabit.quantity, rat.quantity, frog.quantity, rabit.percentage, rat.percentage, frog.percentage)
}
