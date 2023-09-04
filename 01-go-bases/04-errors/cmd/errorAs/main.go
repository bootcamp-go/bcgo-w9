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
	cl := calculator.NewCalculator(0)

	result, err := cl.Divide(n, d)
	if err != nil {
		var ce *calculator.CustomError
		if errors.As(err, &ce) {
			switch ce.Code {
				case 404:
					fmt.Printf("code error 404: %v\n", err)
				default:
					fmt.Printf("code error 500: %v\n", err)
			}
			return
		}
	}

	fmt.Printf("%v / %v = %v\n", n, d, result)
}