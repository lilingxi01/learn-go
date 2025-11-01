package main

import "fmt"

// Solution to Challenge: Personalized Greeting Program
//
// This solution demonstrates:
// - Using multiple fmt.Println() calls
// - Printing strings with proper formatting
// - Adding blank lines for readability

func main() {
	// 1. Print greeting with name
	fmt.Println("Hello, my name is Alex")

	// 2. Print favorite programming language
	fmt.Println("My favorite programming language is Python")

	// 3. Print why learning Go
	fmt.Println("I want to learn Go because it's fast, simple, and great for building scalable backend services")

	// Bonus: You can also print a blank line for formatting
	fmt.Println()

	// Bonus: You can print multiple items in one line
	fmt.Println("Current status:", "Excited to learn Go!")
}

// Alternative approach: You could also use fmt.Printf for formatted output
// We'll learn about fmt.Printf in the next lessons!

// What you learned:
// - package main creates an executable program
// - import "fmt" gives you printing functions
// - func main() is where your program starts
// - fmt.Println() prints text and adds a newline
// - You can call fmt.Println() multiple times
// - Go code is simple and readable
