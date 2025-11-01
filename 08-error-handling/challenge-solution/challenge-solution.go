package main

import (
	"errors"
	"fmt"
	"math"
)

// ===================================
// Custom Error Types
// ===================================

type DivisionByZeroError struct {
	Dividend float64
}

func (e *DivisionByZeroError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by zero", e.Dividend)
}

type InvalidInputError struct {
	Input   string
	Message string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input '%s': %s", e.Input, e.Message)
}

type OutOfRangeError struct {
	Value float64
	Min   float64
	Max   float64
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("value %.2f out of range [%.2f, %.2f]", e.Value, e.Min, e.Max)
}

// Temporary indicates whether the error is temporary
func (e *OutOfRangeError) Temporary() bool {
	return false // Range errors are not temporary
}

type CalculationError struct {
	Operation string
	A, B      float64
	Err       error
}

func (e *CalculationError) Error() string {
	return fmt.Sprintf("calculation error: %.2f %s %.2f - %v", e.A, e.Operation, e.B, e.Err)
}

func (e *CalculationError) Unwrap() error {
	return e.Err
}

// ===================================
// Main Functions
// ===================================

const (
	MinValue = -1000
	MaxValue = 1000
)

func ValidateNumber(n float64) error {
	if math.IsNaN(n) {
		return &InvalidInputError{Input: "NaN", Message: "not a number"}
	}
	if math.IsInf(n, 0) {
		return &InvalidInputError{Input: "Inf", Message: "infinity"}
	}
	if n < MinValue || n > MaxValue {
		return &OutOfRangeError{Value: n, Min: MinValue, Max: MaxValue}
	}
	return nil
}

func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivisionByZeroError{Dividend: a}
	}

	// Validate inputs
	if err := ValidateNumber(a); err != nil {
		return 0, fmt.Errorf("dividend validation failed: %w", err)
	}
	if err := ValidateNumber(b); err != nil {
		return 0, fmt.Errorf("divisor validation failed: %w", err)
	}

	result := a / b
	return result, nil
}

func Calculate(operation string, a, b float64) (float64, error) {
	// Validate inputs
	if err := ValidateNumber(a); err != nil {
		return 0, &CalculationError{Operation: operation, A: a, B: b, Err: err}
	}
	if err := ValidateNumber(b); err != nil {
		return 0, &CalculationError{Operation: operation, A: a, B: b, Err: err}
	}

	var result float64
	var err error

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result, err = SafeDivide(a, b)
		if err != nil {
			return 0, &CalculationError{Operation: operation, A: a, B: b, Err: err}
		}
	default:
		return 0, &InvalidInputError{Input: operation, Message: "unsupported operation"}
	}

	// Validate result
	if err := ValidateNumber(result); err != nil {
		return 0, fmt.Errorf("result validation failed: %w", err)
	}

	return result, nil
}

// Batch processing with error aggregation
func CalculateBatch(operations []Operation) ([]float64, []error) {
	results := make([]float64, len(operations))
	errs := make([]error, 0)

	for i, op := range operations {
		result, err := Calculate(op.Type, op.A, op.B)
		if err != nil {
			errs = append(errs, fmt.Errorf("operation %d (%s): %w", i, op.Type, err))
			results[i] = 0
		} else {
			results[i] = result
		}
	}

	return results, errs
}

type Operation struct {
	Type string
	A, B float64
}

// Safe wrapper with panic recovery
func SafeCalculate(operation string, a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	return Calculate(operation, a, b)
}

// Logging function that handles different error types
func LogError(err error) {
	if err == nil {
		return
	}

	fmt.Println("\nError Details:")
	fmt.Println("──────────────")

	switch e := err.(type) {
	case *DivisionByZeroError:
		fmt.Printf("❌ Division by Zero: Attempted to divide %.2f by zero\n", e.Dividend)
		fmt.Println("   Severity: HIGH")

	case *OutOfRangeError:
		fmt.Printf("❌ Out of Range: Value %.2f not in [%.2f, %.2f]\n", e.Value, e.Min, e.Max)
		fmt.Printf("   Temporary: %t\n", e.Temporary())

	case *InvalidInputError:
		fmt.Printf("❌ Invalid Input: '%s' - %s\n", e.Input, e.Message)
		fmt.Println("   Severity: MEDIUM")

	case *CalculationError:
		fmt.Printf("❌ Calculation Error: %.2f %s %.2f\n", e.A, e.Operation, e.B)
		fmt.Printf("   Underlying error: %v\n", e.Err)

		// Check underlying error type
		var dzErr *DivisionByZeroError
		if errors.As(e.Err, &dzErr) {
			fmt.Println("   Root cause: Division by zero detected")
		}

	default:
		fmt.Printf("❌ Generic Error: %v\n", err)
	}
}

func main() {
	fmt.Println("=== Safe Calculator Challenge Solution ===\n")

	// ===================================
	// 1. Valid Operations
	// ===================================
	fmt.Println("1. Valid operations:")
	operations := []struct {
		op   string
		a, b float64
	}{
		{"+", 10, 5},
		{"-", 20, 8},
		{"*", 6, 7},
		{"/", 100, 4},
	}

	for _, test := range operations {
		result, err := Calculate(test.op, test.a, test.b)
		if err != nil {
			fmt.Printf("   %.0f %s %.0f = Error: %v\n", test.a, test.op, test.b, err)
		} else {
			fmt.Printf("   %.0f %s %.0f = %.2f\n", test.a, test.op, test.b, result)
		}
	}

	// ===================================
	// 2. Division by Zero
	// ===================================
	fmt.Println("\n2. Division by zero:")
	result, err := SafeDivide(10, 0)
	if err != nil {
		LogError(err)
	} else {
		fmt.Printf("   Result: %.2f\n", result)
	}

	// ===================================
	// 3. Out of Range Numbers
	// ===================================
	fmt.Println("\n3. Out of range numbers:")
	result, err = Calculate("+", 2000, 500)
	if err != nil {
		LogError(err)
	}

	// ===================================
	// 4. Invalid Operations
	// ===================================
	fmt.Println("\n4. Invalid operation:")
	result, err = Calculate("%", 10, 3)
	if err != nil {
		LogError(err)
	}

	// ===================================
	// 5. Batch Operations with Errors
	// ===================================
	fmt.Println("\n5. Batch operations:")
	batchOps := []Operation{
		{"+", 10, 5},
		{"/", 20, 0},   // Division by zero
		{"*", 2000, 5}, // Out of range
		{"-", 30, 10},
		{"%", 10, 3}, // Invalid operation
	}

	results, errs := CalculateBatch(batchOps)

	fmt.Println("\nResults:")
	for i, op := range batchOps {
		fmt.Printf("   %.0f %s %.0f = %.2f\n", op.A, op.Type, op.B, results[i])
	}

	if len(errs) > 0 {
		fmt.Printf("\nErrors occurred: %d\n", len(errs))
		for _, err := range errs {
			fmt.Printf("   - %v\n", err)
		}
	}

	// ===================================
	// 6. Panic Recovery
	// ===================================
	fmt.Println("\n\n6. Safe calculation with panic recovery:")
	result, err = SafeCalculate("/", 100, 5)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %.2f\n", result)
	}

	// ===================================
	// 7. Error Wrapping and Unwrapping
	// ===================================
	fmt.Println("\n7. Error wrapping:")
	err = Calculate("/", 10, 0)
	if err != nil {
		fmt.Println("   Top-level error:", err)

		// Unwrap to find root cause
		var dzErr *DivisionByZeroError
		if errors.As(err, &dzErr) {
			fmt.Println("   Root cause found: Division by zero")
		}
	}

	fmt.Println("\n\n=== What You Learned ===")
	fmt.Println("✓ Creating custom error types")
	fmt.Println("✓ Error wrapping with %w")
	fmt.Println("✓ Type checking with errors.As")
	fmt.Println("✓ Error comparison with errors.Is")
	fmt.Println("✓ Panic and recover patterns")
	fmt.Println("✓ Error aggregation")
	fmt.Println("✓ Context-rich error messages")
	fmt.Println("✓ Temporary error interfaces")
}
