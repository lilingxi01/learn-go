package main

import "fmt"

func main() {
	fmt.Println("=== Go Closures Tutorial ===\n")

	// ===================================
	// 1. Basic Closure
	// ===================================
	fmt.Println("1. Basic closure:")
	x := 10
	increment := func() {
		x++ // Closure captures x
		fmt.Printf("   x is now: %d\n", x)
	}
	increment() // x becomes 11
	increment() // x becomes 12
	fmt.Printf("   Final x: %d\n\n", x)

	// ===================================
	// 2. Closure Factory
	// ===================================
	fmt.Println("2. Closure factory (counter):")
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Printf("   Counter1: %d\n", counter1())
	fmt.Printf("   Counter1: %d\n", counter1())
	fmt.Printf("   Counter2: %d\n", counter2())
	fmt.Printf("   Counter1: %d\n\n", counter1())

	// ===================================
	// 3. Closure with Parameters
	// ===================================
	fmt.Println("3. Closure with parameters:")
	addN := makeAdder(10)
	fmt.Printf("   5 + 10 = %d\n", addN(5))
	fmt.Printf("   20 + 10 = %d\n", addN(20))

	add100 := makeAdder(100)
	fmt.Printf("   5 + 100 = %d\n\n", add100(5))

	// ===================================
	// 4. Closure for Accumulation
	// ===================================
	fmt.Println("4. Accumulator closure:")
	acc := makeAccumulator()
	fmt.Printf("   Add 10: Total = %d\n", acc(10))
	fmt.Printf("   Add 20: Total = %d\n", acc(20))
	fmt.Printf("   Add 30: Total = %d\n\n", acc(30))

	// ===================================
	// 5. Closure in Loop (Common Pitfall)
	// ===================================
	fmt.Println("5. Closure in loop:")
	fmt.Println("   Correct way:")
	correctClosureInLoop()

	// ===================================
	// 6. Closure for Configuration
	// ===================================
	fmt.Println("\n6. Closure for configuration:")
	greetEnglish := makeGreeter("Hello")
	greetSpanish := makeGreeter("Hola")
	greetFrench := makeGreeter("Bonjour")

	fmt.Printf("   %s\n", greetEnglish("Alice"))
	fmt.Printf("   %s\n", greetSpanish("Bob"))
	fmt.Printf("   %s\n", greetFrench("Carol"))

	// ===================================
	// 7. Closure for State Management
	// ===================================
	fmt.Println("\n7. Closure for state (account balance):")
	account := makeAccount(1000)

	fmt.Printf("   Deposit $500: %s\n", account.deposit(500))
	fmt.Printf("   Withdraw $200: %s\n", account.withdraw(200))
	fmt.Printf("   Withdraw $2000: %s\n", account.withdraw(2000))
	fmt.Printf("   Balance: $%.2f\n", account.getBalance())

	// ===================================
	// 8. Closure as Callback
	// ===================================
	fmt.Println("\n8. Closure as callback:")
	processNumbers([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Printf("   Processing: %d -> %d\n", n, n*n)
	})

	// ===================================
	// 9. Closure for Memoization
	// ===================================
	fmt.Println("\n9. Memoized fibonacci:")
	fib := memoizedFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("   fib(%d) = %d\n", i, fib(i))
	}

	fmt.Println("\n=== Closures Tutorial Complete! ===")
}

// ===================================
// Closure Function Definitions
// ===================================

// makeCounter creates a counter closure
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// makeAdder creates a closure that adds n to its argument
func makeAdder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// makeAccumulator creates a closure that accumulates values
func makeAccumulator() func(int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}

// Correct way to use closures in loops
func correctClosureInLoop() {
	funcs := make([]func(), 5)

	for i := 0; i < 5; i++ {
		// Capture loop variable by passing it as parameter
		i := i // Create new variable in loop scope
		funcs[i] = func() {
			fmt.Printf("   Function %d\n", i)
		}
	}

	for _, f := range funcs {
		f()
	}
}

// makeGreeter creates a closure for greeting
func makeGreeter(greeting string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

// Account struct using closures for private state
type Account struct {
	deposit    func(float64) string
	withdraw   func(float64) string
	getBalance func() float64
}

// makeAccount creates an account with private balance
func makeAccount(initial float64) Account {
	balance := initial

	return Account{
		deposit: func(amount float64) string {
			if amount > 0 {
				balance += amount
				return fmt.Sprintf("Deposited $%.2f", amount)
			}
			return "Invalid deposit amount"
		},
		withdraw: func(amount float64) string {
			if amount > balance {
				return "Insufficient funds"
			}
			if amount > 0 {
				balance -= amount
				return fmt.Sprintf("Withdrew $%.2f", amount)
			}
			return "Invalid withdrawal amount"
		},
		getBalance: func() float64 {
			return balance
		},
	}
}

// processNumbers applies a function to each number
func processNumbers(numbers []int, fn func(int)) {
	for _, num := range numbers {
		fn(num)
	}
}

// memoizedFibonacci creates a memoized fibonacci closure
func memoizedFibonacci() func(int) int {
	cache := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}

		// Check cache
		if val, ok := cache[n]; ok {
			return val
		}

		// Calculate and cache
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}

	return fib
}
