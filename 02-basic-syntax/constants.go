package main

import "fmt"

// Package-level constants (available throughout the package)
const AppName = "Go Tutorial"
const Version = "1.0.0"

// Multiple constants using const block
const (
	StatusOK       = 200
	StatusNotFound = 404
	StatusError    = 500
)

// Constants with types
const (
	Pi      float64 = 3.14159265359 // π (pi)
	E       float64 = 2.71828182846 // Euler's number
	MaxSize int     = 100
)

// Untyped constants (take the type based on context)
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

func main() {
	fmt.Println("=== Go Constants Tutorial ===\n")

	// ===================================
	// 1. Basic Constants
	// ===================================
	const greeting = "Hello, World!"
	const year = 2024
	fmt.Println("1. Basic constants:")
	fmt.Printf("   Greeting: %s\n", greeting)
	fmt.Printf("   Year: %d\n\n", year)

	// ===================================
	// 2. Package-Level Constants
	// ===================================
	fmt.Println("2. Package-level constants:")
	fmt.Printf("   App: %s v%s\n\n", AppName, Version)

	// ===================================
	// 3. Constant Blocks
	// ===================================
	fmt.Println("3. HTTP Status constants:")
	fmt.Printf("   OK: %d\n", StatusOK)
	fmt.Printf("   Not Found: %d\n", StatusNotFound)
	fmt.Printf("   Error: %d\n\n", StatusError)

	// ===================================
	// 4. Typed Constants
	// ===================================
	fmt.Println("4. Mathematical constants:")
	fmt.Printf("   Pi: %.11f\n", Pi)
	fmt.Printf("   E: %.11f\n\n", E)

	// Using constants in calculations
	radius := 5.0
	circumference := 2 * Pi * radius
	fmt.Printf("   Circle with radius %.1f has circumference %.2f\n\n", radius, circumference)

	// ===================================
	// 5. iota - The Constant Generator
	// ===================================
	fmt.Println("5. Days of the week (using iota):")
	fmt.Printf("   Sunday: %d\n", Sunday)
	fmt.Printf("   Monday: %d\n", Monday)
	fmt.Printf("   Friday: %d\n\n", Friday)

	// iota with expressions
	const (
		_  = iota             // 0 (ignored with _)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
		TB                    // 1 << 40 = 1099511627776
	)

	fmt.Println("6. File size constants (powers of 2):")
	fmt.Printf("   1 KB = %d bytes\n", KB)
	fmt.Printf("   1 MB = %d bytes\n", MB)
	fmt.Printf("   1 GB = %d bytes\n", GB)
	fmt.Printf("   1 TB = %d bytes\n\n", TB)

	// ===================================
	// 7. Constants vs Variables
	// ===================================
	var variable = 10
	variable = 20 // Variables can be changed
	fmt.Println("7. Constants vs Variables:")
	fmt.Printf("   Variable can change: %d\n", variable)
	fmt.Printf("   Constant cannot change: %d\n\n", MaxSize)

	// This would cause a compile error:
	// MaxSize = 200  // Error: cannot assign to MaxSize

	// ===================================
	// 8. Untyped Constants
	// ===================================
	// Untyped constants can be used in any compatible context
	const untypedConst = 42

	var intVar int = untypedConst
	var floatVar float64 = untypedConst
	var complexVar complex128 = untypedConst

	fmt.Println("8. Untyped constant flexibility:")
	fmt.Printf("   As int: %d\n", intVar)
	fmt.Printf("   As float64: %f\n", floatVar)
	fmt.Printf("   As complex128: %v\n\n", complexVar)

	// ===================================
	// 9. Constants in Expressions
	// ===================================
	const (
		width  = 1920
		height = 1080
	)

	area := width * height
	aspectRatio := float64(width) / float64(height)

	fmt.Println("9. Using constants in calculations:")
	fmt.Printf("   Screen: %dx%d\n", width, height)
	fmt.Printf("   Area: %d pixels\n", area)
	fmt.Printf("   Aspect ratio: %.2f\n\n", aspectRatio)

	// ===================================
	// 10. Best Practices
	// ===================================
	fmt.Println("10. Constant Best Practices:")
	fmt.Println("    ✓ Use constants for values that never change")
	fmt.Println("    ✓ Use UPPERCASE or PascalCase for constants")
	fmt.Println("    ✓ Group related constants in const blocks")
	fmt.Println("    ✓ Use iota for sequential values")
	fmt.Println("    ✓ Constants are evaluated at compile time")

	fmt.Println("\n=== Tutorial Complete! ===")
}
