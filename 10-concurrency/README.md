# Concurrency in Go

Master Go's powerful concurrency features with goroutines and channels.

## Prerequisites

- Completed [09-packages-modules](../09-packages-modules/)
- Understanding of functions and interfaces

## What You'll Learn

- Goroutines (lightweight threads)
- Channels for communication
- Select statement for multiple channels
- Sync package (Mutex, WaitGroup, Once)
- Common concurrency patterns
- Race conditions and how to avoid them

## What is Concurrency?

Concurrency is about dealing with multiple things at once. Go makes concurrency a first-class citizen with:

- **Goroutines**: Lightweight threads managed by Go runtime
- **Channels**: Safe communication between goroutines
- **Select**: Multiplexing channel operations

### Concurrency vs Parallelism

- **Concurrency**: Structure of the program (multiple tasks can be in progress)
- **Parallelism**: Execution (multiple tasks running simultaneously)

Go gives you concurrency; the runtime handles parallelism.

## Goroutines

Start a goroutine with the `go` keyword:

```go
func sayHello() {
    fmt.Println("Hello")
}

go sayHello()  // Runs concurrently
```

Goroutines are:

- Lightweight (2KB initial stack)
- Cheap to create (thousands or millions)
- Scheduled by Go runtime

## Channels

Channels enable safe communication between goroutines:

```go
ch := make(chan int)    // Unbuffered channel
ch := make(chan int, 5) // Buffered channel (capacity 5)

ch <- 42      // Send to channel
value := <-ch // Receive from channel
```

### Channel Operations

```go
// Send
ch <- value

// Receive
value := <-ch

// Receive and check if closed
value, ok := <-ch

// Close channel
close(ch)
```

## Select Statement

Select lets you wait on multiple channel operations:

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No activity")
}
```

## Sync Package

### WaitGroup

Wait for goroutines to finish:

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()

wg.Wait()  // Block until all Done()
```

### Mutex

Protect shared data:

```go
var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()
```

### Once

Execute code exactly once:

```go
var once sync.Once

once.Do(func() {
    // This runs only once
})
```

## Running the Examples

```bash
go run goroutines.go
go run channels.go
go run select.go
go run sync.go
```

## Common Patterns

### Worker Pool

```go
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start workers
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
```

### Fan-Out, Fan-In

Multiple goroutines read from same channel (fan-out), results merged into single channel (fan-in).

### Pipeline

Chain of stages connected by channels.

## Best Practices

1. **Don't communicate by sharing memory; share memory by communicating**
2. **Use channels for orchestration; use mutexes for state**
3. **Close channels from sender side only**
4. **Use buffered channels to prevent goroutine leaks**
5. **Always use context for cancellation**

## Common Mistakes

1. **Starting goroutine without waiting**: Use WaitGroup
2. **Closing channel from receiver**: Only sender should close
3. **Not handling channel close**: Check `ok` value
4. **Data races**: Use `-race` flag to detect
5. **Goroutine leaks**: Ensure all goroutines can exit

## Detecting Race Conditions

```bash
go run -race program.go
go test -race ./...
```

## Quick Reference

```go
// Goroutine
go function()

// Channel
ch := make(chan int)
ch <- value    // Send
value := <-ch  // Receive
close(ch)      // Close

// Buffered channel
ch := make(chan int, 100)

// Select
select {
case v := <-ch1:
case ch2 <- v:
case <-time.After(time.Second):
default:
}

// WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
}()
wg.Wait()

// Mutex
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
```

## Next Steps

1. Run the examples
2. Experiment with goroutines and channels
3. Try the worker pool pattern
4. Move to **11-testing** to learn Go testing

## Further Reading

- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency)
- [Share Memory By Communicating](https://go.dev/blog/codelab-share)
