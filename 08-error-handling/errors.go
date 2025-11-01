package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== Go Error Handling Tutorial ===\n")

	// ===================================
	// 1. Basic Error Creation
	// ===================================
	fmt.Println("1. Creating errors:")

	err1 := errors.New("simple error")
	fmt.Printf("   errors.New: %v\n", err1)

	err2 := fmt.Errorf("formatted error: code %d", 404)
	fmt.Printf("   fmt.Errorf: %v\n\n", err2)

	// ===================================
	// 2. Error Handling Pattern
	// ===================================
	fmt.Println("2. Standard error handling:")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n\n", err)
	}

	// ===================================
	// 3. Error Wrapping (Go 1.13+)
	// ===================================
	fmt.Println("3. Error wrapping:")

	err = outer()
	if err != nil {
		fmt.Printf("   Wrapped error: %v\n\n", err)
	}

	// ===================================
	// 4. Error Checking with errors.Is
	// ===================================
	fmt.Println("4. Checking errors with errors.Is:")

	var ErrNotFound = errors.New("not found")

	err = findUser(999)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("   User not found (handled specially)")
	}
	fmt.Println()

	// ===================================
	// 5. Error Type Extraction with errors.As
	// ===================================
	fmt.Println("5. Error type extraction with errors.As:")

	err = validateAge(-5)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("   Validation error - Field: %s, Value: %v\n\n",
			ve.Field, ve.Value)
	}

	// ===================================
	// 6. Multiple Return Values
	// ===================================
	fmt.Println("6. Multiple return values with error:")

	name, age, err := parseUserInput("Alice,25")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Parsed: %s, %d years old\n", name, age)
	}

	_, _, err = parseUserInput("invalid")
	if err != nil {
		fmt.Printf("   Error: %v\n\n", err)
	}

	// ===================================
	// 7. Error Chaining
	// ===================================
	fmt.Println("7. Error chaining:")

	err = level3()
	if err != nil {
		fmt.Printf("   Chained error:\n   %v\n\n", err)

		// Unwrap to see original
		fmt.Println("   Unwrapping:")
		current := err
		for current != nil {
			fmt.Printf("   - %v\n", current)
			current = errors.Unwrap(current)
		}
	}
	fmt.Println()

	// ===================================
	// 8. Panic and Recover
	// ===================================
	fmt.Println("8. Panic and recover:")

	safeFunction()
	fmt.Println("   Program continues after recover\n")

	// ===================================
	// 9. Defer with Error Handling
	// ===================================
	fmt.Println("9. Defer with error handling:")
	processFile("data.txt")
	fmt.Println()

	// ===================================
	// 10. Sentinel Errors
	// ===================================
	fmt.Println("10. Sentinel errors:")

	var (
		ErrInvalidInput = errors.New("invalid input")
		ErrTimeout      = errors.New("timeout")
		ErrUnauthorized = errors.New("unauthorized")
	)

	err = checkAccess("guest")
	switch {
	case errors.Is(err, ErrUnauthorized):
		fmt.Println("    Please log in")
	case errors.Is(err, ErrTimeout):
		fmt.Println("    Request timed out")
	case errors.Is(err, ErrInvalidInput):
		fmt.Println("    Invalid input")
	default:
		if err != nil {
			fmt.Printf("    Unknown error: %v\n", err)
		} else {
			fmt.Println("    Access granted")
		}
	}

	fmt.Println("\n=== Error Handling Tutorial Complete! ===")
}

// ===================================
// Helper Functions
// ===================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func inner() error {
	return errors.New("inner error")
}

func outer() error {
	err := inner()
	if err != nil {
		return fmt.Errorf("outer failed: %w", err)
	}
	return nil
}

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

type ValidationError struct {
	Field string
	Value interface{}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid value %v for field %s", e.Value, e.Field)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Value: age}
	}
	return nil
}

func parseUserInput(input string) (string, int, error) {
	// Simple parser: "name,age"
	var name string
	var age int

	n, err := fmt.Sscanf(input, "%[^,],%d", &name, &age)
	if err != nil || n != 2 {
		return "", 0, fmt.Errorf("invalid format: %s", input)
	}

	return name, age, nil
}

func level3() error {
	return fmt.Errorf("level3: %w", level2())
}

func level2() error {
	return fmt.Errorf("level2: %w", level1())
}

func level1() error {
	return errors.New("level1: original error")
}

func riskyFunction() {
	panic("something went wrong!")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("   Before panic")
	riskyFunction()
	fmt.Println("   After panic (never executes)")
}

func processFile(filename string) {
	defer func() {
		fmt.Println("   Cleanup: closing file (deferred)")
	}()

	fmt.Printf("   Processing file: %s\n", filename)

	if filename == "" {
		fmt.Println("   Error: empty filename")
		return
	}

	fmt.Println("   File processed successfully")
}

func checkAccess(user string) error {
	if user == "guest" {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
