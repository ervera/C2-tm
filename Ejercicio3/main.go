package main

import (
	"encoding/json"
	"fmt"
)

var categories = map[string]float64{"A": 3000, "B": 1500, "C": 1000}

func main() {
	categorie := "A"
	minits := 60

	jsonCategories, _ := json.Marshal(categories)
	fmt.Printf("%v \n", string(jsonCategories))

	fmt.Printf("Categoria: %v  \nMinutos: %v\nSalario: %.2f\n", categorie, minits, salary(categorie, minits))
}

func salary(category string, minits int) (salary float64) {
	value, exists := categories[category]
	if exists {
		salary = (value / 60) * float64(minits)
	}
	return salary
}

/**
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
**/
