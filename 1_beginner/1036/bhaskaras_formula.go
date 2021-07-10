package main

import (
	"fmt"
	"math"
)

type Bhaskara struct {
	termA, termB, termC, r1, r2 float64
}

func getSecondTerm(bhaskara Bhaskara) float64 {
	return math.Pow(bhaskara.termB, 2) - (4 * bhaskara.termA * bhaskara.termC)
}

func isInvalid(bhaskara Bhaskara) bool {
	return bhaskara.termA == 0 || getSecondTerm(bhaskara) < 0
}

func calculate(bhaskara Bhaskara) Bhaskara {
	secondTerm := math.Sqrt(getSecondTerm(bhaskara))
	divisorTerm := 2 * bhaskara.termA

	bhaskara.r1 = (-bhaskara.termB + secondTerm) / divisorTerm
	bhaskara.r2 = (-bhaskara.termB - secondTerm) / divisorTerm

	return bhaskara
}

func main() {
	var bhaskara Bhaskara

	fmt.Scanf("%f %f %f", &bhaskara.termA, &bhaskara.termB, &bhaskara.termC)

	if isInvalid(bhaskara) {
		fmt.Println("Impossivel calcular")

		return
	}

	result := calculate(bhaskara)

	fmt.Printf("R1 = %.5f\nR2 = %.5f\n", result.r1, result.r2)
}
