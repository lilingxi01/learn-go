// Package mathutil provides mathematical utility functions for common operations.
package mathutil

// Add returns the sum of two integers.
// Example: Add(2, 3) returns 5
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
// Example: Multiply(3, 4) returns 12
func Multiply(a, b int) int {
	return a * b
}

// Average calculates the arithmetic mean of a slice of float64 numbers.
// Returns 0.0 for an empty slice.
// Example: Average([]float64{1, 2, 3, 4, 5}) returns 3.0
func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

// validate is unexported (private) - only usable within this package.
// Returns true if n is non-negative.
func validate(n int) bool {
	return n >= 0
}
