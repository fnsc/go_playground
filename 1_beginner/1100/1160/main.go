package main

import "fmt"

type city struct {
	population  int
	growthIndex float64
	years       int
}

func main() {
	entriesQuantity := 0

	read(&entriesQuantity)

	cities := make([][2]city, entriesQuantity)

	for i := 0; i < entriesQuantity; i++ {
		fmt.Scanf("%d", &cities[i][0].population)
		fmt.Scanf("%d", &cities[i][1].population)
		fmt.Scanf("%f", &cities[i][0].growthIndex)
		fmt.Scanf("%f", &cities[i][1].growthIndex)

		calculate(&cities[i])
	}

	printResult(&cities)
}

func read(input *int) {
	fmt.Scanf("%d", input)

	if *input < 1 || *input > 3000 {
		read(input)
	}
}

func calculate(cities *[2]city) {
	for cities[0].population <= cities[1].population && cities[0].years <= 100 {
		cities[0].population += int(float64(cities[0].population) * cities[0].growthIndex / 100)
		cities[1].population += int(float64(cities[1].population) * cities[1].growthIndex / 100)

		cities[0].years++
	}
}

func printResult(cities *[][2]city) {
	for _, city := range *cities {
		if city[0].years > 100 {
			fmt.Println("Mais de 1 seculo.")

			continue
		}

		fmt.Println(city[0].years, "anos.")
	}
}
