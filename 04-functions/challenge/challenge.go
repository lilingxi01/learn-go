package main

// Challenge: Calculator with Error Handling
//
// Create a calculator that performs basic arithmetic operations with proper error handling.
//
// Requirements:
// 1. Create separate functions for: add, subtract, multiply, divide
// 2. Each function should return (result float64, error)
// 3. The divide function must handle division by zero
// 4. Create a calculate() function that takes:
//    - Two numbers (float64)
//    - An operation string ("+", "-", "*", "/")
//    - Returns (result float64, error)
// 5. Test your calculator with:
//    - Valid operations
//    - Division by zero
//    - Invalid operations
//
// Bonus Challenges:
// 6. Add power (exponentiation) operation
// 7. Add modulus operation for integers
// 8. Keep track of calculation history using closures
//
// Hints:
// - Use errors.New() or fmt.Errorf() to create errors
// - Always check errors before using results
// - Use switch statement in calculate() function
//
// Example usage:
//   result, err := calculate(10, 5, "+")
//   if err != nil {
//       fmt.Println("Error:", err)
//       return
//   }
//   fmt.Printf("Result: %.2f\n", result)

func main() {
	// TODO: Implement your calculator here

	// Test cases to implement:
	// 1. 10 + 5 = 15
	// 2. 10 - 5 = 5
	// 3. 10 * 5 = 50
	// 4. 10 / 5 = 2
	// 5. 10 / 0 = error
	// 6. 10 % 3 = invalid operation error
}
