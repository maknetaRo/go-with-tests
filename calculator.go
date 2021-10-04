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


func Subtract(inputs... float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, input := range inputs[1:] {
		result -= input
	}
	return result
	
}

func Multiply(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	var result float64 = inputs[0]
	for _, input := range inputs[1:] {
		result *= input
	}
	return result

}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
		return a / b, nil	
}

func DivideMany(inputs... float64) (float64, error) {
	if len(inputs) == 0 {
		return 0, nil 
	}
	result := inputs[0]
	for _, input := range inputs[1:] {
		if input == 0 {
			return 0, errors.New("division by zero not allowed")
		}
		result = result / input
	}
	return result, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("sqrt of negative number is not a valid operation")
	}
	return math.Sqrt(a), nil
}
