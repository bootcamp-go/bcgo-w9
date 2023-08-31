package main

import "fmt"

func main() {
	n := 1.0
	d := 2.0

	r, err := Divide(n, d)
	if err != "" {
		panic(err)
	}

	fmt.Println(r)
}

func Divide(n, d float64) (float64, string) {
	if d == 0 {
		return 0, "Division by zero is not allowed"
	}

	return n / d, ""
}