package main

import "fmt"

var data [3]int

func main() {
	keepWorking := 0

	for keepWorking < 1 {
		score := readInput()
		store(score)

		action := askForAction()

		if !startOver(action) {
			keepWorking++
		}
	}

	printResult()
}

func readInput() []int {
	counter := 2
	input := 0
	score := make([]int, 0)

	for counter > 0 {
		fmt.Scanf("%d", &input)

		score = append(score, input)
		counter--
	}

	return score
}

func store(score []int) {
	if score[0] > score[1] {
		data[0]++

		return
	}

	if score[1] > score[0] {
		data[1]++

		return
	}

	data[2]++
}

func printResult() {
	total := data[0] + data[1] + data[2]

	if data[0] > data[1] {
		fmt.Printf("%d grenais\nInter:%d\nGremio:%d\nEmpates:%d\nInter venceu mais\n", total, data[0], data[1], data[2])

		return
	}

	if data[0] < data[1] {
		fmt.Printf("%d grenais\nInter:%d\nGremio:%d\nEmpates:%d\nGremio venceu mais\n", total, data[0], data[1], data[2])

		return
	}

	fmt.Printf("%d grenais\nInter:%d\nGremio:%d\nEmpates:%d\nNão houve vencedor\n", total, data[0], data[1], data[2])
}

func askForAction() int {
	input := 0

	fmt.Println("Novo grenal (1-sim 2-nao)")

	fmt.Scanf("%d", &input)

	return input
}

func startOver(action int) bool {
	if action == 1 {
		return true
	}

	if action == 2 {
		return false
	}

	action = askForAction()

	return startOver(action)
}
