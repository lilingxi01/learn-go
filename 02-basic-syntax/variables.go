package main

import "fmt"

func main() {
	fmt.Println("=== Go Variables Tutorial ===\n")

	// ===================================
	// 1. Variable Declaration with var
	// ===================================
	var name string = "Alice"
	var age int = 30
	fmt.Println("1. Explicit declaration:")
	fmt.Printf("   Name: %s, Age: %d\n\n", name, age)

	// ===================================
	// 2. Type Inference (var without type)
	// ===================================
	var city = "San Francisco" // Type inferred as string
	var year = 2024            // Type inferred as int
	fmt.Println("2. Type inference with var:")
	fmt.Printf("   City: %s, Year: %d\n\n", city, year)

	// ===================================
	// 3. Short Declaration (:= operator)
	// ===================================
	// Most common way in Go!
	country := "USA"
	population := 331_000_000 // Underscores for readability
	fmt.Println("3. Short declaration (:=):")
	fmt.Printf("   Country: %s, Population: %d\n\n", country, population)

	// ===================================
	// 4. Multiple Variable Declaration
	// ===================================
	var x, y, z int = 1, 2, 3
	fmt.Println("4. Multiple variables:")
	fmt.Printf("   x=%d, y=%d, z=%d\n\n", x, y, z)

	// Multiple with short declaration
	firstName, lastName := "John", "Doe"
	fmt.Println("5. Multiple with := :")
	fmt.Printf("   Full name: %s %s\n\n", firstName, lastName)

	// ===================================
	// 6. Zero Values
	// ===================================
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string
	fmt.Println("6. Zero values (when not initialized):")
	fmt.Printf("   int: %d\n", defaultInt)           // 0
	fmt.Printf("   float64: %f\n", defaultFloat)     // 0.000000
	fmt.Printf("   bool: %t\n", defaultBool)         // false
	fmt.Printf("   string: '%s'\n\n", defaultString) // ""

	// ===================================
	// 7. Different Numeric Types
	// ===================================
	var smallInt int8 = 127 // -128 to 127
	var largeInt int64 = 9223372036854775807
	var decimal float32 = 3.14
	var precise float64 = 3.141592653589793
	fmt.Println("7. Numeric types:")
	fmt.Printf("   int8: %d\n", smallInt)
	fmt.Printf("   int64: %d\n", largeInt)
	fmt.Printf("   float32: %f\n", decimal)
	fmt.Printf("   float64: %.15f\n\n", precise)

	// ===================================
	// 8. Boolean Type
	// ===================================
	isActive := true
	isComplete := false
	fmt.Println("8. Boolean values:")
	fmt.Printf("   isActive: %t, isComplete: %t\n\n", isActive, isComplete)

	// ===================================
	// 9. String Type
	// ===================================
	message := "Hello, Go!"
	multiline := `This is a
	multi-line string
	using backticks`
	fmt.Println("9. Strings:")
	fmt.Println("   ", message)
	fmt.Println("   ", multiline, "\n")

	// ===================================
	// 10. Type Conversion (must be explicit)
	// ===================================
	var intValue int = 42
	var floatValue float64 = float64(intValue) // Explicit conversion
	var uintValue uint = uint(floatValue)
	fmt.Println("10. Type conversion:")
	fmt.Printf("    int: %d → float64: %f → uint: %d\n\n", intValue, floatValue, uintValue)

	// ===================================
	// 11. Operators
	// ===================================
	a := 10
	b := 3
	fmt.Println("11. Arithmetic operators:")
	fmt.Printf("    %d + %d = %d\n", a, b, a+b)
	fmt.Printf("    %d - %d = %d\n", a, b, a-b)
	fmt.Printf("    %d * %d = %d\n", a, b, a*b)
	fmt.Printf("    %d / %d = %d\n", a, b, a/b)    // Integer division
	fmt.Printf("    %d %% %d = %d\n\n", a, b, a%b) // Modulus

	// Comparison operators
	fmt.Println("12. Comparison operators:")
	fmt.Printf("    10 == 3: %t\n", a == b)
	fmt.Printf("    10 != 3: %t\n", a != b)
	fmt.Printf("    10 > 3: %t\n", a > b)
	fmt.Printf("    10 < 3: %t\n\n", a < b)

	// Logical operators
	fmt.Println("13. Logical operators:")
	fmt.Printf("    true && false: %t\n", true && false)
	fmt.Printf("    true || false: %t\n", true || false)
	fmt.Printf("    !true: %t\n", !true)

	fmt.Println("\n=== Tutorial Complete! ===")
}
