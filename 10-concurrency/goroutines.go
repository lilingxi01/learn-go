package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Goroutines Tutorial ===\n")

	// ===================================
	// 1. Basic Goroutine
	// ===================================
	fmt.Println("1. Basic goroutine:")

	// Regular function call (synchronous)
	sayHello("Alice")

	// Goroutine (asynchronous)
	go sayHello("Bob")

	// Sleep to let goroutine finish
	// (In production, use WaitGroup instead)
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 2. Multiple Goroutines
	// ===================================
	fmt.Println("2. Multiple goroutines:")

	for i := 0; i < 5; i++ {
		go printNumber(i)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 3. Anonymous Function Goroutine
	// ===================================
	fmt.Println("3. Anonymous function goroutine:")

	go func() {
		fmt.Println("   Anonymous goroutine executed")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 4. Goroutine with Parameters
	// ===================================
	fmt.Println("4. Goroutine with parameters:")

	message := "Hello from closure"
	go func(msg string) {
		fmt.Printf("   Received: %s\n", msg)
	}(message)

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 5. Common Mistake - Loop Variable
	// ===================================
	fmt.Println("5. Loop variable capture (common mistake):")

	// ❌ WRONG WAY
	fmt.Println("   Wrong way (all print same value):")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("   Wrong: %d\n", i) // Captures i reference
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// ✅ CORRECT WAY
	fmt.Println("\n   Correct way (pass as parameter):")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("   Correct: %d\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 6. Goroutine Execution Order
	// ===================================
	fmt.Println("6. Execution order (non-deterministic):")

	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("   Goroutine %d\n", id)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 7. Long-Running Goroutine
	// ===================================
	fmt.Println("7. Long-running goroutine:")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("   Working... %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("   Work complete!")
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 8. Goroutine Scheduling
	// ===================================
	fmt.Println("8. CPU-intensive goroutines:")

	go cpuIntensive("A")
	go cpuIntensive("B")

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 9. Main Goroutine Exit
	// ===================================
	fmt.Println("9. Main goroutine behavior:")
	fmt.Println("   Starting goroutines...")

	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("   This might not print if main exits early!")
	}()

	// Main must wait for goroutines
	time.Sleep(100 * time.Millisecond)
	fmt.Println("   Main goroutine waiting...")
	time.Sleep(150 * time.Millisecond)

	fmt.Println("\n=== Goroutines Tutorial Complete! ===")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("✓ Goroutines are lightweight (2KB initial stack)")
	fmt.Println("✓ Use 'go' keyword to start a goroutine")
	fmt.Println("✓ Main goroutine must wait for child goroutines")
	fmt.Println("✓ Execution order is non-deterministic")
	fmt.Println("✓ Always pass loop variables as parameters")
	fmt.Println("✓ In production, use sync.WaitGroup, not time.Sleep")
}

// Helper functions

func sayHello(name string) {
	fmt.Printf("   Hello, %s!\n", name)
}

func printNumber(n int) {
	fmt.Printf("   Number: %d\n", n)
}

func cpuIntensive(id string) {
	// Simulate CPU work
	count := 0
	for i := 0; i < 1000000; i++ {
		count++
	}
	fmt.Printf("   CPU-intensive %s completed: %d iterations\n", id, count)
}
