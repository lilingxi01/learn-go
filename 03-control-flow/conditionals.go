package main

import "fmt"

func main() {
	fmt.Println("=== Go Control Flow: Conditionals ===\n")

	// ===================================
	// 1. Basic If Statement
	// ===================================
	age := 25
	if age >= 18 {
		fmt.Println("1. Basic if: You are an adult")
	}

	// ===================================
	// 2. If-Else
	// ===================================
	temperature := 15
	if temperature > 20 {
		fmt.Println("2. If-else: It's warm outside")
	} else {
		fmt.Println("2. If-else: It's cool outside")
	}

	// ===================================
	// 3. If-Else-If Chain
	// ===================================
	score := 85
	fmt.Print("3. If-else-if: Your grade is ")
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	// ===================================
	// 4. If with Short Statement
	// ===================================
	// Declare variable in the if statement itself
	if num := 42; num%2 == 0 {
		fmt.Printf("4. If with short statement: %d is even\n", num)
	} else {
		fmt.Printf("4. If with short statement: %d is odd\n", num)
	}
	// Note: num is only available within the if/else block

	// ===================================
	// 5. Multiple Conditions
	// ===================================
	userAge := 25
	hasLicense := true

	if userAge >= 18 && hasLicense {
		fmt.Println("5. Multiple conditions: You can drive")
	} else {
		fmt.Println("5. Multiple conditions: You cannot drive")
	}

	// ===================================
	// 6. Nested If Statements
	// ===================================
	isWeekend := true
	isRaining := false

	fmt.Print("6. Nested if: ")
	if isWeekend {
		if isRaining {
			fmt.Println("Stay home and watch movies")
		} else {
			fmt.Println("Go to the park")
		}
	} else {
		fmt.Println("Go to work")
	}

	// ===================================
	// 7. Switch Statement (Basic)
	// ===================================
	day := "Monday"
	fmt.Print("7. Basic switch: ")

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Wednesday":
		fmt.Println("Hump day!")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular weekday")
	}

	// ===================================
	// 8. Switch Without Expression
	// ===================================
	// This is like if-else chains but cleaner
	time := 14
	fmt.Print("8. Switch without expression: ")

	switch {
	case time < 12:
		fmt.Println("Good morning!")
	case time < 17:
		fmt.Println("Good afternoon!")
	case time < 21:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}

	// ===================================
	// 9. Switch with Short Statement
	// ===================================
	switch hour := 15; {
	case hour < 12:
		fmt.Println("9. Switch with short statement: Morning")
	case hour < 18:
		fmt.Println("9. Switch with short statement: Afternoon")
	default:
		fmt.Println("9. Switch with short statement: Evening")
	}

	// ===================================
	// 10. Type Switch (advanced)
	// ===================================
	var value interface{} = 42
	fmt.Print("10. Type switch: Value is a ")

	switch v := value.(type) {
	case int:
		fmt.Printf("int with value %d\n", v)
	case string:
		fmt.Printf("string with value %s\n", v)
	case bool:
		fmt.Printf("bool with value %t\n", v)
	default:
		fmt.Printf("unknown type\n")
	}

	// ===================================
	// 11. Logical Operators in Conditions
	// ===================================
	x := 10
	fmt.Println("\n11. Logical operators:")

	// AND operator
	if x > 5 && x < 15 {
		fmt.Printf("   %d is between 5 and 15 (AND)\n", x)
	}

	// OR operator
	if x < 5 || x > 8 {
		fmt.Printf("   %d is either less than 5 OR greater than 8\n", x)
	}

	// NOT operator
	isEmpty := false
	if !isEmpty {
		fmt.Println("   List is not empty (NOT)")
	}

	// ===================================
	// 12. Comparing Multiple Values
	// ===================================
	a, b, c := 10, 20, 30
	fmt.Println("\n12. Comparing multiple values:")

	if a < b && b < c {
		fmt.Printf("   %d < %d < %d is true\n", a, b, c)
	}

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	fmt.Printf("   Maximum of %d, %d, %d is %d\n", a, b, c, max)

	// ===================================
	// 13. Early Return Pattern (in a function)
	// ===================================
	result := divide(10, 2)
	fmt.Printf("\n13. Early return pattern: 10 / 2 = %.2f\n", result)

	result = divide(10, 0)
	fmt.Printf("    10 / 0 = %.2f (protected by early return)\n", result)

	fmt.Println("\n=== Conditionals Tutorial Complete! ===")
}

// Helper function demonstrating early return pattern
func divide(a, b float64) float64 {
	// Early return for error case
	if b == 0 {
		fmt.Println("    Warning: Division by zero, returning 0")
		return 0
	}
	return a / b
}
