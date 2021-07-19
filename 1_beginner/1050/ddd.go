package main

import "fmt"

type DDD struct {
	city   string
	number int
}

func data() []DDD {
	var brasilia, salvador, saoPaulo, rioDeJaneiro, juizDeFora, campinas, vitoria, beloHorizonte DDD

	brasilia.city = "Brasilia"
	brasilia.number = 61

	salvador.city = "Salvador"
	salvador.number = 71

	saoPaulo.city = "Sao Paulo"
	saoPaulo.number = 11

	rioDeJaneiro.city = "Rio de Janeiro"
	rioDeJaneiro.number = 21

	juizDeFora.city = "Juiz de Fora"
	juizDeFora.number = 32

	campinas.city = "Campinas"
	campinas.number = 19

	vitoria.city = "Vitoria"
	vitoria.number = 27

	beloHorizonte.city = "Belo Horizonte"
	beloHorizonte.number = 31

	return []DDD{brasilia, salvador, saoPaulo, rioDeJaneiro, juizDeFora, campinas, vitoria, beloHorizonte}
}

func search(input int) DDD {
	var ddd DDD
	ddd.city = "DDD nao cadastrado"

	for _, ddd := range data() {
		if !(ddd.number == input) {
			continue
		}

		return ddd
	}

	return ddd
}

func main() {
	var input int

	fmt.Scanf("%d", &input)

	result := search(input)

	fmt.Println(result.city)
}
