package main

import "fmt"

type Animal struct {
	code           int
	name           string
	belongsToOrder int
	belongsToClass int
}

type Class struct {
	code      int
	name      string
	belongsTo int
}

type Order struct {
	code int
	name string
}

type Filo struct {
	code int
	name string
}

func getFilos() []Filo {
	var vertebrate, invertebrate Filo

	vertebrate.code = 1
	vertebrate.name = "vertebrado"

	invertebrate.code = 2
	invertebrate.name = "invertebrado"

	filos := []Filo{vertebrate, invertebrate}

	return filos
}

func getClass() []Class {
	var birds, mammal, insects, annelid Class

	birds.code = 1
	birds.name = "ave"
	birds.belongsTo = 1

	mammal.code = 2
	mammal.name = "mamifero"
	mammal.belongsTo = 1

	insects.code = 3
	insects.name = "inseto"
	insects.belongsTo = 2

	annelid.code = 4
	annelid.name = "anelideo"
	annelid.belongsTo = 2

	return []Class{birds, mammal, insects, annelid}
}

func getOrder() []Order {
	var canivorous, omnivorous, herbivore, hematophagous Order

	canivorous.code = 1
	canivorous.name = "carnivoro"

	omnivorous.code = 2
	omnivorous.name = "onivoro"

	herbivore.code = 3
	herbivore.name = "herbivoro"

	hematophagous.code = 4
	hematophagous.name = "hematofago"

	return []Order{canivorous, omnivorous, herbivore, hematophagous}
}

func getAnimals() []Animal {
	var eagle, pidgeon, man, cown, flea, caterpillar, leech, worm Animal

	eagle.code = 1
	eagle.name = "aguia"
	eagle.belongsToOrder = 1
	eagle.belongsToClass = 1

	pidgeon.code = 2
	pidgeon.name = "pomba"
	pidgeon.belongsToOrder = 2
	pidgeon.belongsToClass = 1

	man.code = 3
	man.name = "homem"
	man.belongsToOrder = 2
	man.belongsToClass = 2

	cown.code = 4
	cown.name = "vaca"
	cown.belongsToOrder = 3
	cown.belongsToClass = 2

	flea.code = 5
	flea.name = "pulga"
	flea.belongsToOrder = 4
	flea.belongsToClass = 3

	caterpillar.code = 6
	caterpillar.name = "lagarta"
	caterpillar.belongsToOrder = 3
	caterpillar.belongsToClass = 3

	leech.code = 7
	leech.name = "sanguessuga"
	leech.belongsToOrder = 4
	leech.belongsToClass = 4

	worm.code = 8
	worm.name = "minhoca"
	worm.belongsToOrder = 2
	worm.belongsToClass = 4

	return []Animal{eagle, pidgeon, man, cown, flea, caterpillar, leech, worm}
}

func searchFilo(input string, table []Filo) Filo {
	var row Filo
	for _, row := range table {
		if !(input == row.name) {
			continue
		}

		return row
	}

	return row
}

func searchOrder(input string, table []Order) Order {
	var row Order
	for _, row := range table {
		if !(input == row.name) {
			continue
		}

		return row
	}

	return row
}

func searchClass(input string, table []Class, code int) Class {
	var row Class
	for _, row := range table {
		if !(input == row.name && code == row.belongsTo) {
			continue
		}

		return row
	}

	return row
}

func searchAnimal(filo Filo, order Order, class Class, table []Animal) Animal {
	var row Animal
	for _, row := range table {
		if !(class.code == row.belongsToClass && order.code == row.belongsToOrder) {
			continue
		}

		return row
	}

	return row
}

func main() {
	var inputFilo, inputClass, inputOrder string

	filos := getFilos()
	orders := getOrder()
	classes := getClass()
	animals := getAnimals()

	fmt.Scanf("%s", &inputFilo)
	fmt.Scanf("%s", &inputClass)
	fmt.Scanf("%s", &inputOrder)

	filo := searchFilo(inputFilo, filos)
	class := searchClass(inputClass, classes, filo.code)
	order := searchOrder(inputOrder, orders)
	animal := searchAnimal(filo, order, class, animals)

	fmt.Println(animal.name)
}
