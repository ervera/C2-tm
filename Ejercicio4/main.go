package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	oper := calc(minimum)
	result, er := oper(1, 2, 3, 4, 5)
	if er == nil {
		fmt.Println(result)
	} else {
		fmt.Println(er)
	}
}

func calc(operation string) func(values ...int) (float64, error) {
	switch operation {
	case minimum:
		return minCalc
	case average:
		return avgCalc
	case maximum:
		return maxCalc
	}
	return nil
}

func minCalc(values ...int) (min float64, er error) {
	min = float64(values[0])
	for _, v := range values {
		if float64(v) <= min {
			min = float64(v)
		}
		if v < 0 {
			return 0, errors.New("Existe al menos un elemento negativo.")
		}
	}
	return min, nil
}

func avgCalc(values ...int) (avg float64, er error) {
	for _, v := range values {
		avg += float64(v)
		if v < 0 {
			return 0, errors.New("Existe al menos un elemento negativo.")
		}
	}
	return avg / float64(len(values)), nil
}

func maxCalc(values ...int) (max float64, er error) {
	for _, v := range values {
		if float64(v) >= max {
			max = float64(v)
		}
		if v < 0 {
			return 0, errors.New("Existe al menos un elemento negativo.")
		}
	}
	return max, nil
}
