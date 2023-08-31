package calculator

const (
	OprAdd = "+"
	OprSub = "-"
	OprMul = "*"
	OprDiv = "/"
)

type OpArithmetic = func(float64, float64) float64

// Add returns the sum of two integers
func Add(a, b float64) (result float64) {
	result = a + b
	return
}

// Subtract returns the difference of two integers
func Subtract(a, b float64) (result float64) {
	result = a - b
	return
}

// Multiply returns the product of two integers
func Multiply(a, b float64) (result float64) {
	result = a * b
	return
}

// Divide returns the quotient of two integers
func Divide(n, d float64) (result float64) {
	if d == 0 {
		return
	}

	result = n / d
	return
}