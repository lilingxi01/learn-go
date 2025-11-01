package main

import "fmt"

func main() {
	fmt.Println("=== Go Control Flow: Loops ===\n")

	// ===================================
	// 1. Traditional For Loop
	// ===================================
	fmt.Println("1. Traditional for loop (0 to 4):")
	for i := 0; i < 5; i++ {
		fmt.Printf("   i = %d\n", i)
	}

	// ===================================
	// 2. While-Style Loop
	// ===================================
	fmt.Println("\n2. While-style loop:")
	count := 0
	for count < 3 {
		fmt.Printf("   count = %d\n", count)
		count++
	}

	// ===================================
	// 3. Infinite Loop with Break
	// ===================================
	fmt.Println("\n3. Infinite loop with break:")
	counter := 0
	for {
		counter++
		if counter > 3 {
			fmt.Println("   Breaking out of infinite loop")
			break
		}
		fmt.Printf("   Iteration %d\n", counter)
	}

	// ===================================
	// 4. Continue Statement
	// ===================================
	fmt.Println("\n4. Continue (skip even numbers):")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("   Odd number: %d\n", i)
	}

	// ===================================
	// 5. Range Over Slice
	// ===================================
	fmt.Println("\n5. Range over slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("   Index %d: %s\n", index, fruit)
	}

	// ===================================
	// 6. Range with Blank Identifier
	// ===================================
	fmt.Println("\n6. Range (ignoring index with _):")
	for _, fruit := range fruits {
		fmt.Printf("   Fruit: %s\n", fruit)
	}

	// ===================================
	// 7. Range Over String (runes)
	// ===================================
	fmt.Println("\n7. Range over string (Unicode runes):")
	text := "Go! 🚀"
	for i, char := range text {
		fmt.Printf("   Position %d: %c (Unicode: %U)\n", i, char, char)
	}

	// ===================================
	// 8. Range Over Map
	// ===================================
	fmt.Println("\n8. Range over map:")
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 28,
	}
	for name, age := range ages {
		fmt.Printf("   %s is %d years old\n", name, age)
	}

	// ===================================
	// 9. Nested Loops
	// ===================================
	fmt.Println("\n9. Nested loops (multiplication table 1-3):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("   %d × %d = %d\n", i, j, i*j)
		}
		if i < 3 {
			fmt.Println()
		}
	}

	// ===================================
	// 10. Loop with Multiple Variables
	// ===================================
	fmt.Println("\n10. Loop with multiple variables:")
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("    i=%d, j=%d\n", i, j)
	}

	// ===================================
	// 11. Range Over Array
	// ===================================
	fmt.Println("\n11. Range over array:")
	numbers := [5]int{10, 20, 30, 40, 50}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("    Array: %v\n", numbers)
	fmt.Printf("    Sum: %d\n", sum)

	// ===================================
	// 12. Breaking Out of Nested Loops
	// ===================================
	fmt.Println("\n12. Breaking nested loops (using label):")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("    i=%d, j=%d\n", i, j)
			if i == 1 && j == 1 {
				fmt.Println("    Breaking out of both loops!")
				break OuterLoop
			}
		}
	}

	// ===================================
	// 13. Iterating Backwards
	// ===================================
	fmt.Println("\n13. Iterating backwards:")
	for i := 5; i >= 1; i-- {
		fmt.Printf("    Countdown: %d\n", i)
	}
	fmt.Println("    Blast off! 🚀")

	// ===================================
	// 14. Skipping Elements
	// ===================================
	fmt.Println("\n14. Processing every 2nd element:")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(data); i += 2 {
		fmt.Printf("    data[%d] = %d\n", i, data[i])
	}

	// ===================================
	// 15. Loop Performance Tips
	// ===================================
	fmt.Println("\n15. Performance tip - cache length:")

	// Good: Cache the length
	largeSlice := make([]int, 1000)
	length := len(largeSlice)
	for i := 0; i < length; i++ {
		largeSlice[i] = i
	}
	fmt.Printf("    Initialized slice of length %d\n", length)

	fmt.Println("\n=== Loops Tutorial Complete! ===")
}
