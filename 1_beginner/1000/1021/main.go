package main

import (
	"fmt"
)

type Money struct {
	quantity int
	value    int
}

func getQuantity(moneys [12]Money, total int) [12]Money {
	var result [12]Money

	for index, money := range moneys {
		money.quantity = total / money.value
		total -= money.value * money.quantity

		result[index] = money

	}

	return result
}

func main() {
	var total float64
	var money [12]Money
	var hundred, fifty, twenty, ten, five, two, one,
		fiftyCents, twentyFiveCents, tenCents, fiveCents, oneCent Money

	hundred.value = 10000
	hundred.quantity = 0

	fifty.value = 5000
	fifty.quantity = 0

	twenty.value = 2000
	twenty.quantity = 0

	ten.value = 1000
	ten.quantity = 0

	five.value = 500
	five.quantity = 0

	two.value = 200
	two.quantity = 0

	one.value = 100
	one.quantity = 0

	fiftyCents.value = 50
	fiftyCents.quantity = 0

	twentyFiveCents.value = 25
	twentyFiveCents.quantity = 0

	tenCents.value = 10
	tenCents.quantity = 0

	fiveCents.value = 5
	fiveCents.quantity = 0

	oneCent.value = 1
	oneCent.quantity = 0

	money[0] = hundred
	money[1] = fifty
	money[2] = twenty
	money[3] = ten
	money[4] = five
	money[5] = two
	money[6] = one
	money[7] = fiftyCents
	money[8] = twentyFiveCents
	money[9] = tenCents
	money[10] = fiveCents
	money[11] = oneCent

	fmt.Scanf("%f", &total)

	result := getQuantity(money, int(total*100))

	fmt.Printf(
		"NOTAS:\n%d nota(s) de R$ 100.00\n%d nota(s) de R$ 50.00\n%d nota(s) de R$ 20.00\n%d nota(s) de R$ 10.00\n%d nota(s) de R$ 5.00\n%d nota(s) de R$ 2.00\nMOEDAS:\n%d moeda(s) de R$ 1.00\n%d moeda(s) de R$ 0.50\n%d moeda(s) de R$ 0.25\n%d moeda(s) de R$ 0.10\n%d moeda(s) de R$ 0.05\n%d moeda(s) de R$ 0.01\n",
		result[0].quantity,
		result[1].quantity,
		result[2].quantity,
		result[3].quantity,
		result[4].quantity,
		result[5].quantity,
		result[6].quantity,
		result[7].quantity,
		result[8].quantity,
		result[9].quantity,
		result[10].quantity,
		result[11].quantity,
	)
}
