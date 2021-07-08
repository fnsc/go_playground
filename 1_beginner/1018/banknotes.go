package main

import (
	"fmt"
)

type Banknote struct {
	quantity, value int
}

func getBanknoteQuantity(banknotes [7]Banknote, total int) [7]Banknote {
	var result [7]Banknote

	for index, banknote := range banknotes {
		banknote.quantity = total / banknote.value
		total = total - (banknote.value * banknote.quantity)

		result[index] = banknote
	}

	return result
}

func main() {
	var total int
	var banknotes [7]Banknote
	var hundred, fifty, twenty, ten, five, two, one Banknote

	hundred.value = 100
	hundred.quantity = 0

	fifty.value = 50
	fifty.quantity = 0

	twenty.value = 20
	twenty.quantity = 0

	ten.value = 10
	ten.quantity = 0

	five.value = 5
	five.quantity = 0

	two.value = 2
	two.quantity = 0

	one.value = 1
	one.quantity = 0

	banknotes[0] = hundred
	banknotes[1] = fifty
	banknotes[2] = twenty
	banknotes[3] = ten
	banknotes[4] = five
	banknotes[5] = two
	banknotes[6] = one

	fmt.Scanf("%d", &total)

	result := getBanknoteQuantity(banknotes, total)

	fmt.Printf(
		"%d\n%d nota(s) de R$ 100,00\n%d nota(s) de R$ 50,00\n%d nota(s) de R$ 20,00\n%d nota(s) de R$ 10,00\n%d nota(s) de R$ 5,00\n%d nota(s) de R$ 2,00\n%d nota(s) de R$ 1,00\n",
		total,
		result[0].quantity,
		result[1].quantity,
		result[2].quantity,
		result[3].quantity,
		result[4].quantity,
		result[5].quantity,
		result[6].quantity,
	)
}
