package calculator

import (
	"errors"
	"fmt"
)

// Custom Error
func NewCustomError(Code int, Op string, MaxUsage int, Msg string) *CustomError {
	return &CustomError{
		Code: Code,
		Op: Op,
		MaxUsage: MaxUsage,
		Msg: Msg,
	}
}
type CustomError struct {
	Code int
	Op string
	MaxUsage int
	Msg string
}
func (e *CustomError) Error() string {
	msg := fmt.Sprintf("op: %v, max usage: %v, msg: %v", e.Op, e.MaxUsage, e.Msg)
	return msg
}

// _________________________________________________________________
var (
	ErrDivideByZero = errors.New("cannot divide by zero")
	ErrMaxUsage = errors.New("max usage reached")
	ErrInternal = errors.New("internal error")

	// ErrCustom01 = NewCustomError(401, "divide", 1, "max usage reached")
)

func NewCalculator(maxUsage int) *Calculator {
	return &Calculator{
		maxUsage: maxUsage,
	}
}

type Calculator struct {
	maxUsage int
	currentUsage int
}

// Add takes two floats and returns the result of adding them together.
func (c *Calculator) Add(a, b float64) (result float64) {
	result = a + b
	return
}

// Subtract takes two floats and returns the result of subtracting the second
// from the first.
func (c *Calculator) Subtract(a, b float64) (result float64) {
	result = a - b
	return
}

// Multiply takes two floats and returns the result of multiplying them
// together.
func (c *Calculator) Multiply(a, b float64) (result float64) {
	result = a * b
	return
}

// Divide takes two floats and returns the result of dividing the first by the
// second.
func (c *Calculator) Divide(n, d float64) (result float64, err error) {
	// check available usage
	if c.currentUsage >= c.maxUsage {
		err = NewCustomError(404, "divide", c.maxUsage, "max usage reached")
		return
	}

	// check if denominator is zero
	if d == 0 {
		err = fmt.Errorf("%w. numerator: %v, denominator: %v", ErrDivideByZero, n, d)
		return
	}

	result = n / d
	c.currentUsage++
	return
}