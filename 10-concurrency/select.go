package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Select Statement Tutorial ===\n")

	// ===================================
	// 1. Basic Select
	// ===================================
	fmt.Println("1. Basic select with two channels:")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("   Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("   Received: %s\n", msg2)
		}
	}
	fmt.Println()

	// ===================================
	// 2. Select with Default
	// ===================================
	fmt.Println("2. Select with default (non-blocking):")

	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Printf("   Received: %s\n", msg)
	default:
		fmt.Println("   No message available (default case)")
	}
	fmt.Println()

	// ===================================
	// 3. Select with Timeout
	// ===================================
	fmt.Println("3. Select with timeout:")

	result := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		result <- "operation completed"
	}()

	select {
	case res := <-result:
		fmt.Printf("   %s\n", res)
	case <-time.After(1 * time.Second):
		fmt.Println("   Timeout: operation took too long")
	}
	fmt.Println()

	// ===================================
	// 4. Select for Cancellation
	// ===================================
	fmt.Println("4. Select for cancellation:")

	done := make(chan bool)
	data := make(chan int)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("   Worker: Received cancellation signal")
				return
			case num := <-data:
				fmt.Printf("   Worker: Processing %d\n", num)
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		data <- i
		time.Sleep(50 * time.Millisecond)
	}

	done <- true
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 5. Select with Multiple Receivers
	// ===================================
	fmt.Println("5. Select with multiple receivers:")

	source1 := make(chan string)
	source2 := make(chan string)
	source3 := make(chan string)

	go func() {
		source1 <- "message from source 1"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		source2 <- "message from source 2"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		source3 <- "message from source 3"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg := <-source1:
			fmt.Printf("   From source1: %s\n", msg)
		case msg := <-source2:
			fmt.Printf("   From source2: %s\n", msg)
		case msg := <-source3:
			fmt.Printf("   From source3: %s\n", msg)
		}
	}
	fmt.Println()

	// ===================================
	// 6. Random Selection
	// ===================================
	fmt.Println("6. Random selection when multiple ready:")

	ready1 := make(chan string, 1)
	ready2 := make(chan string, 1)

	ready1 <- "message 1"
	ready2 <- "message 2"

	// Both channels are ready, Go picks randomly
	select {
	case msg := <-ready1:
		fmt.Printf("   Selected: %s\n", msg)
	case msg := <-ready2:
		fmt.Printf("   Selected: %s\n", msg)
	}
	fmt.Println()

	// ===================================
	// 7. Tick Pattern
	// ===================================
	fmt.Println("7. Ticker pattern with select:")

	ticker := time.NewTicker(100 * time.Millisecond)
	done2 := make(chan bool)

	go func() {
		time.Sleep(400 * time.Millisecond)
		done2 <- true
	}()

	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("   Tick %d\n", count)
		case <-done2:
			ticker.Stop()
			fmt.Println("   Ticker stopped")
			fmt.Println()
			return
		}
	}
}
