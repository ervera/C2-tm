package main

import (
	"errors"
	"fmt"
)

//var foodPerAnimal = map[string]int{"perro": 10000, "gato": 5000, "hamster": 250, "tarantula": 150}

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func main() {
	animalName := cat
	count := 100
	operation := animal(animalName)
	txt, er := operation(count)
	if er == nil {
		fmt.Printf("El animal %v con %v cantidad necesita de %v gramos de alimento. \n", animalName, count, txt)
	} else {
		fmt.Println(er)
	}
}

func animal(animalType string) func(count int) (int, error) {
	switch animalType {
	case dog:
		return dogCal
	case cat:
		return catCal
	case hamster:
		return hamsterCal
	case spider:
		return spiderCal
	default:
		return animalNoExist
	}
}

func dogCal(count int) (textResult int, er error) {
	return (count * 10000), er
}

func catCal(count int) (textResult int, er error) {
	return (count * 5000), er
}

func hamsterCal(count int) (textResult int, er error) {
	return (count * 250), er
}

func spiderCal(count int) (textResult int, er error) {
	return (count * 150), er
}

func animalNoExist(count int) (textResult int, er error) {
	return (count * 150), errors.New("el animal no existe")
}

/** MODULO 5 EJECERCICIO 5
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

**/
