package main

import "fmt"

// Solution: Temperature Converter
//
// This solution demonstrates:
// - Variable declaration with type inference
// - Arithmetic operations
// - Type consistency (all float64)
// - Formatted output with fmt.Printf

func main() {
	fmt.Println("=== Temperature Converter ===\n")

	// ===================================
	// Basic Solution
	// ===================================

	// Step 1: Declare temperature in Celsius
	celsius := 25.0

	// Step 2: Convert to Fahrenheit
	// Formula: F = C × 9/5 + 32
	const celsiusToFahrenheitRatio = 9.0 / 5.0
	const fahrenheitOffset = 32.0
	fahrenheit := celsius*celsiusToFahrenheitRatio + fahrenheitOffset

	// Step 3: Convert to Kelvin
	// Formula: K = C + 273.15
	const absoluteZeroCelsius = 273.15
	kelvin := celsius + absoluteZeroCelsius

	// Step 4: Print all values with 2 decimal places
	fmt.Println("Basic Conversion:")
	fmt.Printf("Temperature: %.2f°C\n", celsius)
	fmt.Printf("Temperature: %.2f°F\n", fahrenheit)
	fmt.Printf("Temperature: %.2fK\n\n", kelvin)

	// ===================================
	// Bonus: Multiple Temperature Tests
	// ===================================

	fmt.Println("Testing Different Temperatures:")
	fmt.Println("--------------------------------")

	// Test with different temperatures
	temperatures := []float64{-40.0, 0.0, 25.0, 100.0}

	for _, c := range temperatures {
		f := c*celsiusToFahrenheitRatio + fahrenheitOffset
		k := c + absoluteZeroCelsius

		fmt.Printf("\n%.2f°C = %.2f°F = %.2fK\n", c, f, k)
	}

	// ===================================
	// Bonus: Interesting Temperature Facts
	// ===================================

	fmt.Println("\n\nInteresting Temperature Facts:")
	fmt.Println("--------------------------------")

	// -40°C = -40°F (same in both scales!)
	const specialTempCelsius = -40.0
	specialF := specialTempCelsius*celsiusToFahrenheitRatio + fahrenheitOffset
	fmt.Printf("-40°C = %.0f°F (they meet at -40!)\n", specialF)

	// Water freezing point
	const waterFreezingC = 0.0
	freezingF := waterFreezingC*celsiusToFahrenheitRatio + fahrenheitOffset
	freezingK := waterFreezingC + absoluteZeroCelsius
	fmt.Printf("\nWater freezes at: %.0f°C = %.0f°F = %.2fK\n",
		waterFreezingC, freezingF, freezingK)

	// Water boiling point
	const waterBoilingC = 100.0
	boilingF := waterBoilingC*celsiusToFahrenheitRatio + fahrenheitOffset
	boilingK := waterBoilingC + absoluteZeroCelsius
	fmt.Printf("Water boils at: %.0f°C = %.0f°F = %.2fK\n",
		waterBoilingC, boilingF, boilingK)

	// ===================================
	// Alternative Approach with Constants
	// ===================================

	fmt.Println("\n\nUsing Constants:")
	fmt.Println("--------------------------------")

	// Define conversion factors as constants
	const (
		CelsiusToFahrenheitRatio = 9.0 / 5.0
		FahrenheitOffset         = 32.0
		CelsiusToKelvinOffset    = 273.15
	)

	roomTemp := 22.0
	roomTempF := roomTemp*CelsiusToFahrenheitRatio + FahrenheitOffset
	roomTempK := roomTemp + CelsiusToKelvinOffset

	fmt.Printf("Room temperature: %.1f°C = %.1f°F = %.2fK\n",
		roomTemp, roomTempF, roomTempK)

	// ===================================
	// What You Learned
	// ===================================

	fmt.Println("\n\n=== What You Learned ===")
	fmt.Println("✓ Variable declaration with type inference (var and :=)")
	fmt.Println("✓ Using float64 for decimal numbers")
	fmt.Println("✓ Arithmetic operations (+, -, *, /)")
	fmt.Println("✓ Formatted output with fmt.Printf and %.2f")
	fmt.Println("✓ Using constants for magic numbers")
	fmt.Println("✓ Order of operations in expressions")
}

// Key Takeaways:
// 1. Always use consistent types - all temperatures are float64
// 2. Use meaningful variable names (celsius, fahrenheit, kelvin)
// 3. Constants make formulas more readable
// 4. fmt.Printf with %.2f formats floats to 2 decimal places
// 5. You can test your code with multiple values
