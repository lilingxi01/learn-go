package main

// Challenge: Safe Division Calculator with Comprehensive Error Handling
//
// Build a robust calculator with proper error handling.
//
// Requirements:
// 1. Create custom error types:
//    - DivisionByZeroError
//    - InvalidInputError
//    - OutOfRangeError
//
// 2. Implement functions:
//    - SafeDivide(a, b float64) (float64, error)
//    - ValidateNumber(n float64) error
//    - Calculate(operation string, a, b float64) (float64, error)
//
// 3. Validation rules:
//    - No division by zero
//    - Numbers must be between -1000 and 1000
//    - Support operations: +, -, *, /
//
// 4. Demonstrate error wrapping
// 5. Use panic/recover for critical errors
//
// Bonus Challenges:
// 6. Implement error aggregation for batch operations
// 7. Add Temporary() method to appropriate errors
// 8. Create a logging function that handles different error types
//
// Hints:
// - Use fmt.Errorf with %w for wrapping
// - Check error types with errors.As
// - Recover from panics in deferred functions

func main() {
	// TODO: Implement your solution here

	// Test cases:
	// 1. Valid operations
	// 2. Division by zero
	// 3. Out of range numbers
	// 4. Invalid operations
	// 5. Batch operations with multiple errors
}
