package calculator

import "fmt"

// ===================================
// Example Tests (Testable Documentation)
// ===================================

// ExampleAdd demonstrates the Add function
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

// ExampleSubtract demonstrates the Subtract function
func ExampleSubtract() {
	result := Subtract(10, 3)
	fmt.Println(result)
	// Output: 7
}

// ExampleMultiply demonstrates the Multiply function
func ExampleMultiply() {
	result := Multiply(4, 5)
	fmt.Println(result)
	// Output: 20
}

// ExampleDivide demonstrates the Divide function
func ExampleDivide() {
	result, _ := Divide(10, 2) // Safe: we know divisor is not zero
	fmt.Println(result)
	// Output: 5
}

// ExampleDivide_divisionByZero demonstrates error handling
func ExampleDivide_divisionByZero() {
	_, err := Divide(10, 0)
	fmt.Println(err)
	// Output: division by zero
}

// ExampleSum demonstrates the Sum function
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Println(result)
	// Output: 15
}

// ExampleAverage demonstrates the Average function
func ExampleAverage() {
	numbers := []int{10, 20, 30, 40, 50}
	result := Average(numbers)
	fmt.Printf("%.1f\n", result)
	// Output: 30.0
}

// ExampleIsEven demonstrates the IsEven function
func ExampleIsEven() {
	fmt.Println(IsEven(4))
	fmt.Println(IsEven(5))
	// Output:
	// true
	// false
}

// ExampleAbs demonstrates the Abs function
func ExampleAbs() {
	fmt.Println(Abs(-42))
	fmt.Println(Abs(42))
	// Output:
	// 42
	// 42
}

// ExampleMax demonstrates the Max function
func ExampleMax() {
	result := Max(10, 20)
	fmt.Println(result)
	// Output: 20
}

// ExampleFactorial demonstrates the Factorial function
func ExampleFactorial() {
	result, _ := Factorial(5) // Safe: we know 5 is non-negative
	fmt.Println(result)
	// Output: 120
}
