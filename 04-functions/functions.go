package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== Go Functions Tutorial ===\n")

	// ===================================
	// 1. Basic Function
	// ===================================
	greeting := greet("Alice")
	fmt.Println("1. Basic function:", greeting)

	// ===================================
	// 2. Function with Multiple Parameters
	// ===================================
	sum := add(10, 20)
	fmt.Printf("2. Multiple parameters: 10 + 20 = %d\n", sum)

	// ===================================
	// 3. Multiple Return Values
	// ===================================
	quotient, remainder := divmod(17, 5)
	fmt.Printf("3. Multiple returns: 17 ÷ 5 = %d remainder %d\n", quotient, remainder)

	// ===================================
	// 4. Named Return Values
	// ===================================
	min, max := minMax(5, 10)
	fmt.Printf("4. Named returns: min=%d, max=%d\n", min, max)

	// ===================================
	// 5. Function Returning Error
	// ===================================
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("5. Error:", err)
	} else {
		fmt.Printf("5. Safe divide: 10 / 2 = %.2f\n", result)
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("5. Safe divide error: %v\n", err)
	}

	// ===================================
	// 6. Variadic Functions
	// ===================================
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("\n6. Variadic function: sum(1,2,3,4,5) = %d\n", total)

	numbers := []int{10, 20, 30}
	total = sumAll(numbers...) // Spread operator
	fmt.Printf("6. With spread: sum([10,20,30]) = %d\n", total)

	// ===================================
	// 7. Functions as Values
	// ===================================
	fmt.Println("\n7. Functions as values:")
	operation := add
	fmt.Printf("   Using add function: %d\n", operation(5, 3))

	operation = multiply
	fmt.Printf("   Using multiply function: %d\n", operation(5, 3))

	// ===================================
	// 8. Anonymous Functions
	// ===================================
	fmt.Println("\n8. Anonymous functions:")

	// Immediate execution
	func() {
		fmt.Println("   Immediately executed anonymous function")
	}()

	// Assigned to variable
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("   Square of 5: %d\n", square(5))

	// ===================================
	// 9. Higher-Order Functions
	// ===================================
	fmt.Println("\n9. Higher-order functions:")
	nums := []int{1, 2, 3, 4, 5}
	doubled := mapInts(nums, func(x int) int { return x * 2 })
	fmt.Printf("   Original: %v\n", nums)
	fmt.Printf("   Doubled: %v\n", doubled)

	// ===================================
	// 10. Defer Statement
	// ===================================
	fmt.Println("\n10. Defer statement:")
	demonstrateDefer()

	// ===================================
	// 11. Multiple Defers (LIFO)
	// ===================================
	fmt.Println("\n11. Multiple defers (Last In, First Out):")
	multipleDefers()

	// ===================================
	// 12. Defer with Values
	// ===================================
	fmt.Println("\n12. Defer with parameters:")
	deferWithParams()

	// ===================================
	// 13. Recursive Function
	// ===================================
	fmt.Println("\n13. Recursive function:")
	fmt.Printf("    Factorial of 5: %d\n", factorial(5))
	fmt.Printf("    Fibonacci(10): %d\n", fibonacci(10))

	fmt.Println("\n=== Functions Tutorial Complete! ===")
}

// ===================================
// Function Definitions
// ===================================

// Basic function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple parameters of same type
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// Multiple return values
func divmod(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Named return values
func minMax(a, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return // naked return
}

// Function with error handling
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Variadic function (variable number of arguments)
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Higher-order function (function that takes function as parameter)
func mapInts(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

// Demonstrating defer
func demonstrateDefer() {
	defer fmt.Println("    This prints last (deferred)")
	fmt.Println("    This prints first")
	fmt.Println("    This prints second")
}

// Multiple defers execute in LIFO order
func multipleDefers() {
	defer fmt.Println("    Defer 1 (executed 3rd)")
	defer fmt.Println("    Defer 2 (executed 2nd)")
	defer fmt.Println("    Defer 3 (executed 1st)")
	fmt.Println("    Regular statement (executed first)")
}

// Defer with parameters (evaluated immediately)
func deferWithParams() {
	x := 10
	defer fmt.Printf("    Deferred: x = %d (captured at defer time)\n", x)
	x = 20
	fmt.Printf("    Regular: x = %d\n", x)
}

// Recursive function - Factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Recursive function - Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
