package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Channels Tutorial ===\n")

	// ===================================
	// 1. Unbuffered Channel
	// ===================================
	fmt.Println("1. Unbuffered channel (blocks until receive):")

	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Printf("   Received: %s\n\n", msg)

	// ===================================
	// 2. Buffered Channel
	// ===================================
	fmt.Println("2. Buffered channel (capacity 3):")

	buffered := make(chan int, 3)

	// Can send without blocking up to capacity
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("   Sent 3 values (non-blocking)")

	fmt.Printf("   Received: %d\n", <-buffered)
	fmt.Printf("   Received: %d\n", <-buffered)
	fmt.Printf("   Received: %d\n\n", <-buffered)

	// ===================================
	// 3. Channel Direction (Send/Receive Only)
	// ===================================
	fmt.Println("3. Channel direction:")

	data := make(chan int)
	go sendOnly(data)
	receiveOnly(data)
	fmt.Println()

	// ===================================
	// 4. Closing Channels
	// ===================================
	fmt.Println("4. Closing channels:")

	numbers := make(chan int, 5)

	// Send values
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // Close when done sending
	}()

	// Receive until closed
	for num := range numbers {
		fmt.Printf("   Received: %d\n", num)
	}
	fmt.Println()

	// ===================================
	// 5. Checking if Channel is Closed
	// ===================================
	fmt.Println("5. Checking if channel is closed:")

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	for {
		value, ok := <-ch2
		if !ok {
			fmt.Println("   Channel closed")
			break
		}
		fmt.Printf("   Received: %d\n", value)
	}
	fmt.Println()

	// ===================================
	// 6. Fan-Out Pattern
	// ===================================
	fmt.Println("6. Fan-out pattern (multiple workers):")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
	fmt.Println()

	// ===================================
	// 7. Pipeline Pattern
	// ===================================
	fmt.Println("7. Pipeline pattern:")

	// Stage 1: Generate numbers
	nums := generate(2, 3, 4, 5)

	// Stage 2: Square the numbers
	squares := square(nums)

	// Stage 3: Print results
	for result := range squares {
		fmt.Printf("   Result: %d\n", result)
	}
	fmt.Println()

	// ===================================
	// 8. Timeout with Channels
	// ===================================
	fmt.Println("8. Channel with timeout:")

	slow := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		slow <- "completed"
	}()

	select {
	case msg := <-slow:
		fmt.Printf("   Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("   Timeout! Operation took too long")
	}
	fmt.Println()

	// ===================================
	// 9. Non-Blocking Channel Operations
	// ===================================
	fmt.Println("9. Non-blocking operations with default:")

	messages := make(chan string, 1)

	// Non-blocking send
	select {
	case messages <- "buffered":
		fmt.Println("   Sent message")
	default:
		fmt.Println("   No message sent")
	}

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Printf("   Received: %s\n", msg)
	default:
		fmt.Println("   No message received")
	}
	fmt.Println()

	// ===================================
	// 10. Nil Channel Behavior
	// ===================================
	fmt.Println("10. Nil channel behavior:")

	var nilCh chan int

	select {
	case <-nilCh:
		fmt.Println("   Never executes (nil channel blocks forever)")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("   Nil channel blocks forever, timeout triggered")
	}

	fmt.Println("\n=== Channels Tutorial Complete! ===")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("✓ Unbuffered channels block until both send and receive ready")
	fmt.Println("✓ Buffered channels allow n sends without blocking")
	fmt.Println("✓ Close channels from sender side only")
	fmt.Println("✓ Use range to receive until channel is closed")
	fmt.Println("✓ Check ok value to detect closed channels")
	fmt.Println("✓ Never close a channel from receiver side")
}

// Helper functions

// sendOnly can only send to the channel
func sendOnly(ch chan<- int) {
	ch <- 42
}

// receiveOnly can only receive from the channel
func receiveOnly(ch <-chan int) {
	value := <-ch
	fmt.Printf("   Received via typed channel: %d\n", value)
}

// worker processes jobs from jobs channel and sends results to results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
		results <- job * 2
	}
}

// generate creates a channel and sends values
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square receives from in, squares values, sends to out
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
