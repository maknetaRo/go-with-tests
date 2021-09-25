package calculator

import (
	"errors"
	"math"
)

func Add(inputs  ...float64) float64 {
	var a float64
	for _, input := range inputs {		
		a += input
	}
	return a
}


func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
		return a / b, nil	
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("sqrt of negative number is not a valid operation")
	}
	return math.Sqrt(a), nil
}
