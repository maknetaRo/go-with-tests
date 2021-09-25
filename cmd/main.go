package main 

import (
	"calculator"
	"fmt"
)

func main() {
	result := calculator.Add(-2, 2, 2)
	fmt.Println(result)
	result = calculator.Add(5, 5, 10, 5)
	fmt.Println(result)
}