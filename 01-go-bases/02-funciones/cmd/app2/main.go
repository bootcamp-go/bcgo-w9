package main

import (
	"app/calculator"
	"fmt"
)

func main() {
	n1, n2 := 1.0, 2.0

	// Add
	operationSum := operationArithmetic(calculator.OprAdd)
	fmt.Println(operationSum(n1, n2))

	// Subtract
	operationSub := operationArithmetic(calculator.OprSub)
	fmt.Println(operationSub(n1, n2))

	// Multiply
	operationMul := operationArithmetic(calculator.OprMul)
	fmt.Println(operationMul(n1, n2))

	// Divide
	operationDiv := operationArithmetic(calculator.OprDiv)
	fmt.Println(operationDiv(n1, n2))
}

// operationArithmetic is a function that takes two float64 values and returns a float64 value
func operationArithmetic(operator string) calculator.OpArithmetic {
	switch operator {
	case calculator.OprAdd:
		return calculator.Add
	case calculator.OprSub:
		return calculator.Subtract
	case calculator.OprMul:
		return calculator.Multiply
	case calculator.OprDiv:
		return calculator.Divide
	default:
		return nil
	}
}

// decorator pattern
func decorator(operation calculator.OpArithmetic) (operationDecorated calculator.OpArithmetic) {
	operationDecorated = func(f1, f2 float64) float64 {
		result := operation(f1, f2)
		fmt.Println("Operation result", result)
		return result
	}
	return
}