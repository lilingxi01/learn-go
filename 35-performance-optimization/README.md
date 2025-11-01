# Performance Optimization in Go

Profile and optimize Go applications for production.

## Prerequisites

- Completed [34-docker-containerization](../34-docker-containerization/)

## What You'll Learn

- CPU profiling with pprof
- Memory profiling
- Benchmark writing
- Performance best practices
- Common optimizations
- Identifying bottlenecks

## Profiling Tools

### CPU Profiling

```go
import _ "net/http/pprof"

go func() {
    http.ListenAndServe("localhost:6060", nil)
}()
```

View at: http://localhost:6060/debug/pprof/

### Memory Profiling

```bash
go test -memprofile=mem.out
go tool pprof mem.out
```

### Benchmarking

```go
func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Function()
    }
}
```

Run: `go test -bench=. -benchmem`

## Common Optimizations

1. **Pre-allocate slices**: `make([]T, 0, size)`
2. **Use strings.Builder**: For string concatenation
3. **Pool objects**: `sync.Pool`
4. **Avoid allocations**: Return pointers carefully
5. **Use buffered I/O**: For file operations

## Running

```bash
cd 35-performance-optimization
go run -cpuprofile=cpu.out main.go
go tool pprof cpu.out
```

## Next Steps

Move to **36-production-patterns** for final production patterns

## Further Reading

- [Profiling Go Programs](https://go.dev/blog/pprof)
- [High Performance Go](https://dave.cheney.net/high-performance-go-workshop/gopherchina-2019.html)

