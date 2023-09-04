package main

import (
	"errors"
	"fmt"
)

func main() {
	errBase := errors.New("base error")
	errWrap1 := fmt.Errorf("wrap 1: %w", errBase)

	errWrap2 := fmt.Errorf("wrap 2: %w", errWrap1)

	errWrap3 := fmt.Errorf("wrap 3: %w", errWrap2)

	fmt.Println(errWrap3)

	errUnwrap2 := errors.Unwrap(errWrap3)
	fmt.Println(errUnwrap2)

	errUnwrap1 := errors.Unwrap(errUnwrap2)
	fmt.Println(errUnwrap1)

	errUnwrapBase := errors.Unwrap(errUnwrap1)
	fmt.Println(errUnwrapBase)

	errUnwrapNil := errors.Unwrap(errUnwrapBase)
	fmt.Println(errUnwrapNil)
}