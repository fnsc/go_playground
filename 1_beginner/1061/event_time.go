package main

import (
	"fmt"
	"time"
)

type DateTime struct {
	days, hours, minutes, seconds int
}

func printResult(result DateTime) {
	fmt.Printf("%d dia(s)\n%d hora(s)\n%d minuto(s)\n%d segundo(s)\n", result.days, result.hours, result.minutes, result.seconds)
}

func calculate(seconds float64) {
	var duration DateTime

	duration.days = int(seconds / 60 / 60 / 24)
	duration.hours = int(seconds/60/60) - (duration.days * 24)
	duration.minutes = int(seconds/60) - (duration.days * 24 * 60) - (duration.hours * 60)
	duration.seconds = int(seconds) - (duration.days * 24 * 60 * 60) - (duration.hours * 60 * 60) - (duration.minutes * 60)

	printResult(duration)
}

func main() {
	var inputStart, inputEnd DateTime
	var discart string

	fmt.Scanf("%4s%d", &discart, &inputStart.days)
	fmt.Scanf("%d%3s%d%3s%d", &inputStart.hours, &discart, &inputStart.minutes, &discart, &inputStart.seconds)
	fmt.Scanf("%4s%d", &discart, &inputEnd.days)
	fmt.Scanf("%d%3s%d%3s%d", &inputEnd.hours, &discart, &inputEnd.minutes, &discart, &inputEnd.seconds)

	startDate := time.Date(2021, 4, inputStart.days, inputStart.hours, inputStart.minutes, inputStart.seconds, 0, time.UTC)
	endDate := time.Date(2021, 4, inputEnd.days, inputEnd.hours, inputEnd.minutes, inputEnd.seconds, 0, time.UTC)

	calculate(endDate.Sub(startDate).Seconds())
}
