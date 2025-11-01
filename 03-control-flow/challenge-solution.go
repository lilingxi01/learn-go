package main

import "fmt"

// Solution: Enhanced FizzBuzz
//
// This solution demonstrates multiple approaches to the FizzBuzz problem

func main() {
	fmt.Println("=== Enhanced FizzBuzz Solution ===\n")

	// ===================================
	// Approach 1: String Building
	// ===================================
	fmt.Println("Approach 1: String Building")
	fmt.Println("----------------------------")

	fizzCount, buzzCount, boomCount := 0, 0, 0

	for i := 1; i <= 100; i++ {
		output := ""

		// Check each divisor and build the string
		if i%3 == 0 {
			output += "Fizz"
			fizzCount++
		}
		if i%5 == 0 {
			output += "Buzz"
			buzzCount++
		}
		if i%7 == 0 {
			output += "Boom"
			boomCount++
		}

		// If no special case, print the number
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}

		// Print every 10 numbers on one line for compact display
		if i%10 != 0 {
			fmt.Printf("%-12s ", output)
		} else {
			fmt.Printf("%-12s\n", output)
		}
	}

	fmt.Println("\nStatistics:")
	fmt.Printf("Fizz appeared: %d times\n", fizzCount)
	fmt.Printf("Buzz appeared: %d times\n", buzzCount)
	fmt.Printf("Boom appeared: %d times\n", boomCount)

	// ===================================
	// Approach 2: If-Else Chain
	// ===================================
	fmt.Println("\n\nApproach 2: Selected Examples Using If-Else")
	fmt.Println("--------------------------------------------")

	testCases := []int{3, 5, 7, 15, 21, 35, 105}
	for _, num := range testCases {
		if num%3 == 0 && num%5 == 0 && num%7 == 0 {
			fmt.Printf("%3d: FizzBuzzBoom\n", num)
		} else if num%3 == 0 && num%5 == 0 {
			fmt.Printf("%3d: FizzBuzz\n", num)
		} else if num%3 == 0 && num%7 == 0 {
			fmt.Printf("%3d: FizzBoom\n", num)
		} else if num%5 == 0 && num%7 == 0 {
			fmt.Printf("%3d: BuzzBoom\n", num)
		} else if num%3 == 0 {
			fmt.Printf("%3d: Fizz\n", num)
		} else if num%5 == 0 {
			fmt.Printf("%3d: Buzz\n", num)
		} else if num%7 == 0 {
			fmt.Printf("%3d: Boom\n", num)
		} else {
			fmt.Printf("%3d: %d\n", num, num)
		}
	}

	// ===================================
	// Approach 3: Using Functions
	// ===================================
	fmt.Println("\n\nApproach 3: Interesting Numbers (Using Function)")
	fmt.Println("------------------------------------------------")

	interestingNumbers := []int{15, 21, 30, 35, 42, 70, 105}
	for _, num := range interestingNumbers {
		fmt.Printf("%3d: %s\n", num, fizzBuzzBoom(num))
	}

	// ===================================
	// Approach 4: Complete Range with Analysis
	// ===================================
	fmt.Println("\n\nApproach 4: Analysis of 1-50")
	fmt.Println("-----------------------------")

	regularCount := 0
	specialCount := 0

	for i := 1; i <= 50; i++ {
		result := fizzBuzzBoom(i)
		if result == fmt.Sprintf("%d", i) {
			regularCount++
		} else {
			specialCount++
			fmt.Printf("%3d: %s\n", i, result)
		}
	}

	fmt.Printf("\nIn range 1-50:")
	fmt.Printf("\n  Regular numbers: %d", regularCount)
	fmt.Printf("\n  Special numbers: %d\n", specialCount)

	// ===================================
	// What You Learned
	// ===================================
	fmt.Println("\n=== What You Learned ===")
	fmt.Println("✓ For loops in Go")
	fmt.Println("✓ Modulo operator (%) for divisibility")
	fmt.Println("✓ String concatenation")
	fmt.Println("✓ If-else chains and conditions")
	fmt.Println("✓ Multiple conditions with && (AND)")
	fmt.Println("✓ Building strings dynamically")
	fmt.Println("✓ Formatting output with Printf")
}

// fizzBuzzBoom returns the FizzBuzzBoom value for a number
func fizzBuzzBoom(n int) string {
	result := ""

	if n%3 == 0 {
		result += "Fizz"
	}
	if n%5 == 0 {
		result += "Buzz"
	}
	if n%7 == 0 {
		result += "Boom"
	}

	// If no special case, return the number as a string
	if result == "" {
		return fmt.Sprintf("%d", n)
	}

	return result
}

// Alternative solution using switch (less common for FizzBuzz)
func fizzBuzzBoomSwitch(n int) string {
	div3 := n%3 == 0
	div5 := n%5 == 0
	div7 := n%7 == 0

	switch {
	case div3 && div5 && div7:
		return "FizzBuzzBoom"
	case div3 && div5:
		return "FizzBuzz"
	case div3 && div7:
		return "FizzBoom"
	case div5 && div7:
		return "BuzzBoom"
	case div3:
		return "Fizz"
	case div5:
		return "Buzz"
	case div7:
		return "Boom"
	default:
		return fmt.Sprintf("%d", n)
	}
}
