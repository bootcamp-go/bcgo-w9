package main

import (
	"app/calculator"
	"fmt"
)

func main() {
	// Add
	fmt.Println(orchestrator(calculator.OprAdd, 1, 2, 3, 4, 5))
	
	// Subtract
	fmt.Println(orchestrator(calculator.OprSub, 1, 2, 3, 4, 5))

	// Multiply
	fmt.Println(orchestrator(calculator.OprMul, 1, 2, 3, 4, 5))

	// Divide
	fmt.Println(orchestrator(calculator.OprDiv, 1, 2, 3, 4, 5))
}

func orchestrator(operator string, values ...float64) float64 {
	switch operator {
	case calculator.OprAdd:
		return iterator(calculator.Add, values...)
	case calculator.OprSub:
		return iterator(calculator.Subtract, values...)
	case calculator.OprMul:
		return iterator(calculator.Multiply, values...)
	case calculator.OprDiv:
		return iterator(calculator.Divide, values...)
	default:
		return 0
	}
}

func iterator(operation func(float64, float64) float64, values ...float64) float64 {
	// check if there are at least two values
	if len(values) < 2  {
		return 0
	}

	// Initialize the result with the first value
	result := values[0]

	// Iterate over the remaining values
	for _, value := range values[1:] {
		result = operation(result, value)
	}

	return result
}