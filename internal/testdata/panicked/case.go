// Package panicked has tests with panic at runtime.
package panicked

import "errors"

type TheGood struct{}

func (TheGood) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

type TheBad struct{}

func (TheBad) Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

type TheUgly struct{}

func (TheUgly) Divide(a, b int) int {
	return a / b
}
