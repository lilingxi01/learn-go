package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Calculator Challenge Solution ===\n")

	// ===================================
	// Basic Operations
	// ===================================
	fmt.Println("Basic Operations:")
	fmt.Println("-----------------")

	testCalculations := []struct {
		a, b      float64
		operation string
	}{
		{10, 5, "+"},
		{10, 5, "-"},
		{10, 5, "*"},
		{10, 5, "/"},
		{10, 0, "/"}, // Division by zero
		{10, 3, "^"}, // Power
		{10, 5, "%"}, // Invalid
	}

	for _, test := range testCalculations {
		result, err := calculate(test.a, test.b, test.operation)
		if err != nil {
			fmt.Printf("%.0f %s %.0f = Error: %v\n",
				test.a, test.operation, test.b, err)
		} else {
			fmt.Printf("%.0f %s %.0f = %.2f\n",
				test.a, test.operation, test.b, result)
		}
	}

	// ===================================
	// Using Individual Functions
	// ===================================
	fmt.Println("\nUsing Individual Functions:")
	fmt.Println("---------------------------")

	if result, err := add(15, 25); err == nil {
		fmt.Printf("15 + 25 = %.2f\n", result)
	}

	if result, err := divide(100, 4); err == nil {
		fmt.Printf("100 / 4 = %.2f\n", result)
	}

	if _, err := divide(10, 0); err != nil {
		fmt.Printf("10 / 0 = Error: %v\n", err)
	}

	// ===================================
	// Calculator with History (Bonus)
	// ===================================
	fmt.Println("\nCalculator with History:")
	fmt.Println("------------------------")

	calc := makeCalculatorWithHistory()

	calc.compute(10, 5, "+")
	calc.compute(20, 4, "*")
	calc.compute(100, 10, "/")
	calc.compute(50, 25, "-")

	calc.printHistory()

	// ===================================
	// Advanced Calculator
	// ===================================
	fmt.Println("\nAdvanced Operations:")
	fmt.Println("--------------------")

	operations := []struct {
		desc string
		fn   func() (float64, error)
	}{
		{"Power: 2^8", func() (float64, error) { return power(2, 8) }},
		{"Square root: √16", func() (float64, error) { return sqrt(16) }},
		{"Square root: √-4", func() (float64, error) { return sqrt(-4) }},
		{"Percentage: 20% of 150", func() (float64, error) { return percentage(20, 150) }},
	}

	for _, op := range operations {
		result, err := op.fn()
		if err != nil {
			fmt.Printf("%s = Error: %v\n", op.desc, err)
		} else {
			fmt.Printf("%s = %.2f\n", op.desc, result)
		}
	}

	// ===================================
	// Chain Calculations
	// ===================================
	fmt.Println("\nChained Calculations:")
	fmt.Println("---------------------")

	// ((10 + 5) * 2) / 3
	// Note: In production, check each error. For demonstration, we know these won't fail.
	step1, _ := add(10, 5)         // 15 (error impossible here)
	step2, _ := multiply(step1, 2) // 30 (error impossible here)
	step3, _ := divide(step2, 3)   // 10 (safe division)
	fmt.Printf("((10 + 5) * 2) / 3 = %.2f\n", step3)

	fmt.Println("\n=== What You Learned ===")
	fmt.Println("✓ Functions with multiple return values")
	fmt.Println("✓ Error handling pattern")
	fmt.Println("✓ Using errors.New() and fmt.Errorf()")
	fmt.Println("✓ Switch statements for operation selection")
	fmt.Println("✓ Closures for maintaining state")
	fmt.Println("✓ Function composition and chaining")
}

// ===================================
// Basic Operation Functions
// ===================================

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// ===================================
// Main Calculate Function
// ===================================

func calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return add(a, b)
	case "-":
		return subtract(a, b)
	case "*":
		return multiply(a, b)
	case "/":
		return divide(a, b)
	case "^":
		return power(a, b)
	default:
		return 0, fmt.Errorf("invalid operation: %s", operation)
	}
}

// ===================================
// Advanced Operations
// ===================================

func power(base, exponent float64) (float64, error) {
	return math.Pow(base, exponent), nil
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	return math.Sqrt(n), nil
}

func percentage(percent, total float64) (float64, error) {
	if percent < 0 || percent > 100 {
		return 0, errors.New("percentage must be between 0 and 100")
	}
	return (percent / 100) * total, nil
}

// ===================================
// Calculator with History (Closure)
// ===================================

type CalculatorWithHistory struct {
	compute      func(float64, float64, string)
	printHistory func()
}

func makeCalculatorWithHistory() CalculatorWithHistory {
	history := []string{}

	return CalculatorWithHistory{
		compute: func(a, b float64, op string) {
			result, err := calculate(a, b, op)
			var entry string
			if err != nil {
				entry = fmt.Sprintf("%.0f %s %.0f = Error: %v", a, op, b, err)
			} else {
				entry = fmt.Sprintf("%.0f %s %.0f = %.2f", a, op, b, result)
			}
			history = append(history, entry)
			fmt.Println(entry)
		},
		printHistory: func() {
			fmt.Println("\nCalculation History:")
			for i, entry := range history {
				fmt.Printf("%d. %s\n", i+1, entry)
			}
		},
	}
}
