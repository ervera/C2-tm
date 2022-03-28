package main

import (
	"errors"
	"fmt"
)

func main() {
	res, er := sum(1, 2, 3, 4, 5, 6, 7, 8)

	if er == nil {
		fmt.Printf("La suma total es %d \n", res)
	} else {
		fmt.Println(er)
	}
}

func sum(values ...int) (int, error) {
	var result int
	for _, v := range values {
		if v >= 0 {
			result += v
		} else {
			return 0, errors.New("Existe al menos un elemento negativo.")
		}
	}
	return result, nil
}

/** Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativ
**/
