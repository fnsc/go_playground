package main

import "fmt"

type Snack struct {
	code, value int
	name        string
}

func calculate(code int, amount int) float64 {
	var hotDog, xSalad, xBacon, simpleToast, soda Snack
	var snacks [5]Snack
	var total float64 = 0

	hotDog.code = 1
	hotDog.name = "Cachorro Quente"
	hotDog.value = 400

	xSalad.code = 2
	xSalad.name = "X-Salada"
	xSalad.value = 450

	xBacon.code = 3
	xBacon.name = "X-Bacon"
	xBacon.value = 500

	simpleToast.code = 4
	simpleToast.name = "Torrada Simples"
	simpleToast.value = 200

	soda.code = 5
	soda.name = "Refrigerante"
	soda.value = 150

	snacks[0] = hotDog
	snacks[1] = xSalad
	snacks[2] = xBacon
	snacks[3] = simpleToast
	snacks[4] = soda

	for _, snack := range snacks {
		if !(code == snack.code) {
			continue
		}

		total = float64(snack.value) * float64(amount) / 100

		return total
	}

	return total
}

func main() {
	var code, amount int

	fmt.Scanf("%d %d", &code, &amount)

	result := calculate(code, amount)

	fmt.Printf("Total: R$ %.2f\n", result)
}
