package main

import (
	"app/internal/calculator"
	"errors"
	"fmt"
)

func main() {
	// operations
	n := 10.0
	d := 0.0

	// calculator
	cl := calculator.NewCalculator(2)

	result, err := cl.Divide(n, d)
	if err != nil {
		// if errors.Is(err, calculator.ErrCustom01) {
		// 	errVerify.Code = 1
		// 	fmt.Printf("code error 01: %v\n", err)
		// }
		switch {
			case errors.Is(err, calculator.ErrDivideByZero):
				fmt.Printf("code error 01: %v\n", err)
			case errors.Is(err, calculator.ErrMaxUsage):
				fmt.Printf("code error 02: %v\n", err)
			default:
				fmt.Printf("code error unknown: %v\n", err)
			}
		return
	}

	fmt.Printf("%v / %v = %v\n", n, d, result)
}