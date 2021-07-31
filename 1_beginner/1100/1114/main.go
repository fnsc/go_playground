package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	var password string

	for 0 < 1 {
		fmt.Scanf("%s", &password)

		if isValid(password) {
			printResult("Acesso Permitido")

			return
		}

		printResult("Senha Invalida")
	}
}

func isValid(password string) bool {
	return password == "2002"
}

func printResult(message string) {
	fmt.Println(message)
}
