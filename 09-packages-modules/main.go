package main

import (
	"example.com/packages-tutorial/mathutil"
	"example.com/packages-tutorial/stringutil"
	"fmt"
)

func main() {
	fmt.Println("=== Go Packages and Modules Tutorial ===\n")

	// ===================================
	// 1. Using Custom mathutil Package
	// ===================================
	fmt.Println("1. Using mathutil package:")

	sum := mathutil.Add(10, 5)
	fmt.Printf("   Add(10, 5) = %d\n", sum)

	product := mathutil.Multiply(10, 5)
	fmt.Printf("   Multiply(10, 5) = %d\n", product)

	avg := mathutil.Average([]float64{10, 20, 30, 40, 50})
	fmt.Printf("   Average([10,20,30,40,50]) = %.2f\n\n", avg)

	// ===================================
	// 2. Using Custom stringutil Package
	// ===================================
	fmt.Println("2. Using stringutil package:")

	reversed := stringutil.Reverse("Hello, Go!")
	fmt.Printf("   Reverse('Hello, Go!') = '%s'\n", reversed)

	upper := stringutil.ToUpper("hello world")
	fmt.Printf("   ToUpper('hello world') = '%s'\n", upper)

	isPal := stringutil.IsPalindrome("racecar")
	fmt.Printf("   IsPalindrome('racecar') = %t\n\n", isPal)

	// ===================================
	// 3. Package Visibility
	// ===================================
	fmt.Println("3. Package visibility:")
	fmt.Println("   ✓ Can access: mathutil.Add (exported)")
	fmt.Println("   ✗ Cannot access: mathutil.validate (unexported)")
	fmt.Println("   Exported names start with capital letter")

	fmt.Println("\n=== Tutorial Complete! ===")
	fmt.Println("\nWhat you learned:")
	fmt.Println("✓ Creating packages")
	fmt.Println("✓ Importing local packages")
	fmt.Println("✓ Package visibility rules")
	fmt.Println("✓ Using go.mod for modules")
}
