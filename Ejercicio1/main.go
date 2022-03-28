package main

import (
	"fmt"
)

func main() {
	var salary float64 = 50000
	fmt.Printf("El salario es de %2f \n", getSalary(salary))
}

func getSalary(salary float64) float64 {

	if salary > 50000 {
		salary *= (1 - 0.17)
	} else if salary > 150000 {
		salary *= (1 - 0.1)
	}
	return salary
}

/**
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
**/
