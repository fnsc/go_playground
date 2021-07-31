package main

import (
	"fmt"
)

type GameDuration struct {
	startHour, startMinute,
	endHour, endMinute int
}

func printResult(resultHour, resultMinute int) {
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", resultHour, resultMinute)
}

func main() {
	var game GameDuration

	fmt.Scanf("%d %d %d %d", &game.startHour, &game.startMinute, &game.endHour, &game.endMinute)

	if game.startHour == game.endHour && game.startMinute == game.endMinute {
		fmt.Println("O JOGO DUROU 24 HORA(S) E 0 MINUTO(S)")

		return
	}

	if game.startHour == game.endHour && game.endMinute > game.startMinute {
		resultMinute := game.endMinute - game.startMinute
		resultHour := game.startHour - game.endHour

		printResult(resultHour, resultMinute)

		return
	}

	if game.startHour == game.endHour && game.startMinute > game.endMinute {
		resultMinute := 60 - game.startMinute + game.endMinute
		resultHour := 24 - game.startHour + game.endHour - 1

		printResult(resultHour, resultMinute)

		return
	}

	if game.startMinute == game.endMinute && game.startHour < game.endHour {
		resultMinute := 0
		resultHour := game.endHour - game.startHour

		printResult(resultHour, resultMinute)

		return
	}

	if game.startMinute == game.endMinute && game.startHour > game.endHour {
		resultMinute := 0
		resultHour := 24 - game.startHour + game.endHour

		printResult(resultHour, resultMinute)

		return
	}

	if game.endHour > game.startHour && game.endMinute > game.startMinute {
		resultMinute := game.endMinute - game.startMinute
		resultHour := game.endHour - game.startHour

		printResult(resultHour, resultMinute)

		return
	}

	if game.startHour < game.endHour && game.startMinute > game.endMinute {
		resultMinute := 60 - game.startMinute + game.endMinute
		resultHour := game.endHour - game.startHour - 1

		printResult(resultHour, resultMinute)

		return
	}

	if game.startHour > game.endHour && game.startMinute < game.endMinute {
		resultMinute := game.endMinute - game.startMinute
		resultHour := 24 - game.startHour - 1 + game.endHour

		printResult(resultHour, resultMinute)

		return
	}

	resultMinute := 60 + game.endMinute - game.startMinute
	resultHour := 24 + game.endHour - game.startHour - 1

	printResult(resultHour, resultMinute)

}
