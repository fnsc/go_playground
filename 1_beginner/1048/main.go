package main

import "fmt"

type Salary struct {
	amount, bonusAmount float64
	bonusPercentage     int
}

func calculate(input float64, bonus int) Salary {
	var salary Salary

	salary.bonusAmount = input * (float64(bonus) / 100)
	salary.amount = salary.bonusAmount + input
	salary.bonusPercentage = bonus

	return salary
}

func pickStrategy(input float64) Salary {
	if input > 0 && input <= 400 {
		return calculate(input, 15)
	}

	if input >= 400.01 && input <= 800 {
		return calculate(input, 12)
	}

	if input >= 800.01 && input <= 1200 {
		return calculate(input, 10)
	}

	if input >= 1200.01 && input <= 2000 {
		return calculate(input, 7)
	}

	return calculate(input, 4)
}

func print(salary Salary) {
	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %d %%\n", salary.amount, salary.bonusAmount, salary.bonusPercentage)
}

func main() {
	var input float64

	fmt.Scanf("%f", &input)

	result := pickStrategy(input)

	print(result)
}
