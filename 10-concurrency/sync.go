package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go Sync Package Tutorial ===\n")

	// ===================================
	// 1. WaitGroup - Basic Usage
	// ===================================
	fmt.Println("1. WaitGroup - waiting for goroutines:")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter

		go func(id int) {
			defer wg.Done() // Decrement when done

			fmt.Printf("   Worker %d starting\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("   Worker %d done\n", id)
		}(i)
	}

	wg.Wait() // Block until counter is 0
	fmt.Println("   All workers completed\n")

	// ===================================
	// 2. Mutex - Protecting Shared State
	// ===================================
	fmt.Println("2. Mutex - protecting shared counter:")

	var mu sync.Mutex
	var counter int
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg2.Wait()
	fmt.Printf("   Final counter: %d\n\n", counter)

	// ===================================
	// 3. RWMutex - Readers-Writer Lock
	// ===================================
	fmt.Println("3. RWMutex - multiple readers, single writer:")

	var rwMu sync.RWMutex
	var data = make(map[string]int)
	data["key"] = 0

	// Multiple readers
	for i := 0; i < 3; i++ {
		go func(id int) {
			rwMu.RLock()
			val := data["key"]
			fmt.Printf("   Reader %d: %d\n", id, val)
			rwMu.RUnlock()
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// Single writer
	rwMu.Lock()
	data["key"] = 100
	fmt.Println("   Writer: updated to 100")
	rwMu.Unlock()

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 4. Once - Execute Exactly Once
	// ===================================
	fmt.Println("4. Once - guarantee single execution:")

	var once sync.Once
	var wg3 sync.WaitGroup

	initialize := func() {
		fmt.Println("   Initialization running...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("   Initialization complete")
	}

	// Try to call multiple times
	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			fmt.Printf("   Goroutine %d calling once.Do\n", id)
			once.Do(initialize) // Only runs once
		}(i)
	}

	wg3.Wait()
	fmt.Println()

	// ===================================
	// 5. Cond - Condition Variable
	// ===================================
	fmt.Println("5. Cond - condition variable:")

	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	ready := false

	// Waiter goroutine
	go func() {
		mutex.Lock()
		for !ready {
			cond.Wait() // Releases lock and waits
		}
		fmt.Println("   Waiter: Condition met!")
		mutex.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)

	// Signal goroutine
	mutex.Lock()
	ready = true
	cond.Signal() // Wake up one waiter
	mutex.Unlock()

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// ===================================
	// 6. Pool - Object Reuse
	// ===================================
	fmt.Println("6. Pool - object reuse:")

	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("   Creating new object")
			return &struct{ value int }{value: 0}
		},
	}

	// Get from pool
	obj1 := pool.Get().(*struct{ value int })
	obj1.value = 42
	fmt.Printf("   Got object: %d\n", obj1.value)

	// Put back
	pool.Put(obj1)
	fmt.Println("   Returned object to pool")

	// Get again (might be same object)
	obj2 := pool.Get().(*struct{ value int })
	fmt.Printf("   Got object: %d\n\n", obj2.value)

	// ===================================
	// 7. Map - Concurrent Safe Map
	// ===================================
	fmt.Println("7. sync.Map - concurrent-safe map:")

	var sm sync.Map

	// Store values
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	// Load value
	if val, ok := sm.Load("key1"); ok {
		fmt.Printf("   Found: %s\n", val)
	}

	// LoadOrStore
	actual, loaded := sm.LoadOrStore("key3", "value3")
	fmt.Printf("   LoadOrStore: %s, existed: %t\n", actual, loaded)

	// Range over map
	fmt.Println("   All entries:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("      %s: %s\n", key, value)
		return true // continue iteration
	})
	fmt.Println()

	// ===================================
	// 8. Race Condition Demo
	// ===================================
	fmt.Println("8. Race condition (without mutex) vs safe (with mutex):")

	// Unsafe counter
	unsafeCounter := 0
	var wg4 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			unsafeCounter++ // RACE CONDITION!
		}()
	}
	wg4.Wait()
	fmt.Printf("   Unsafe counter: %d (should be 100, may be less)\n", unsafeCounter)

	// Safe counter
	safeCounter := 0
	var safeMu sync.Mutex
	var wg5 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg5.Add(1)
		go func() {
			defer wg5.Done()
			safeMu.Lock()
			safeCounter++
			safeMu.Unlock()
		}()
	}
	wg5.Wait()
	fmt.Printf("   Safe counter: %d (always 100)\n", safeCounter)

	fmt.Println("\n=== Sync Package Tutorial Complete! ===")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("✓ WaitGroup: Wait for multiple goroutines to complete")
	fmt.Println("✓ Mutex: Protect shared data from concurrent access")
	fmt.Println("✓ RWMutex: Allow multiple readers or single writer")
	fmt.Println("✓ Once: Guarantee code runs exactly once")
	fmt.Println("✓ Pool: Reuse objects to reduce allocations")
	fmt.Println("✓ Map: Thread-safe map for concurrent access")
	fmt.Println("✓ Always use sync primitives to avoid race conditions")
}
