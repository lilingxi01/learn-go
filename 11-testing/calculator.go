// Package calculator provides basic arithmetic operations
package calculator

import (
	"errors"
	"fmt"
)

// ErrDivisionByZero is returned when dividing by zero
var ErrDivisionByZero = errors.New("division by zero")

// Add returns the sum of two integers.
// Example: Add(2, 3) returns 5
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
// Example: Subtract(5, 3) returns 2
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
// Example: Multiply(3, 4) returns 12
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers.
// Returns ErrDivisionByZero if b is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// IsEven returns true if the number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Abs returns the absolute value of an integer
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Max returns the larger of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Sum returns the sum of a slice of integers
func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Average returns the average of a slice of integers.
// Returns 0 for empty slice.
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(Sum(numbers)) / float64(len(numbers))
}

// Factorial calculates n! (factorial of n).
// Returns error for negative numbers.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial of negative number: %d", n)
	}
	if n == 0 || n == 1 {
		return 1, nil
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
