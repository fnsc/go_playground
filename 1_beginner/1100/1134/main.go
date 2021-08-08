package main

import "fmt"

var data [3]int

func main() {
	keepWorking := 0

	for keepWorking < 1 {
		fuel := readInput()

		if !startOver(fuel) {
			keepWorking++
		}

		if fuel >= 1 && fuel <= 3 {
			store(fuel)
		}
	}

	printResult()
}

func readInput() int {
	fuel := 0

	fmt.Scanf("%d", &fuel)

	return fuel
}

func store(fuel int) {
	if fuel == 1 {
		data[0]++

		return
	}

	if fuel == 2 {
		data[1]++

		return
	}

	data[2]++
}

func printResult() {
	fmt.Printf("MUITO OBRIGADO\nAlcool: %d\nGasolina: %d\nDiesel: %d\n", data[0], data[1], data[2])
}

func startOver(action int) bool {
	return action != 4
}
