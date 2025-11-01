# Production Patterns

Essential patterns for production-ready Go services.

## Prerequisites

- Completed [35-performance-optimization](../35-performance-optimization/)

## What You'll Learn

- Graceful shutdown
- Health checks and readiness probes
- Rate limiting
- Circuit breakers
- Retries and timeouts
- Distributed tracing basics
- Metrics and monitoring

## Graceful Shutdown

```go
func main() {
    srv := &http.Server{Addr: ":8080"}
    
    go func() {
        srv.ListenAndServe()
    }()
    
    // Wait for interrupt
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Shutdown error:", err)
    }
}
```

## Health Checks

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
    })
}

func readinessHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := db.Ping(); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}
```

## Rate Limiting

```go
import "golang.org/x/time/rate"

limiter := rate.NewLimiter(10, 100) // 10 req/sec, burst 100

func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

## Circuit Breaker

```go
type CircuitBreaker struct {
    maxFailures int
    resetTimeout time.Duration
    failures    int
    lastFailTime time.Time
    state       string // closed, open, half-open
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    if cb.state == "open" {
        if time.Since(cb.lastFailTime) > cb.resetTimeout {
            cb.state = "half-open"
        } else {
            return errors.New("circuit breaker open")
        }
    }
    
    err := fn()
    if err != nil {
        cb.failures++
        cb.lastFailTime = time.Now()
        if cb.failures >= cb.maxFailures {
            cb.state = "open"
        }
        return err
    }
    
    cb.failures = 0
    cb.state = "closed"
    return nil
}
```

## Retries with Exponential Backoff

```go
func retryWithBackoff(fn func() error, maxRetries int) error {
    var err error
    for i := 0; i < maxRetries; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        
        // Exponential backoff
        waitTime := time.Duration(math.Pow(2, float64(i))) * time.Second
        time.Sleep(waitTime)
    }
    return fmt.Errorf("max retries exceeded: %w", err)
}
```

## Timeouts

```go
func callWithTimeout(url string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    return nil
}
```

## Running

```bash
cd 36-production-patterns
go run main.go
```

## Best Practices Checklist

- [ ] Graceful shutdown implemented
- [ ] Health and readiness endpoints
- [ ] Rate limiting on public APIs
- [ ] Circuit breakers for external services
- [ ] Retries with exponential backoff
- [ ] All operations have timeouts
- [ ] Structured logging
- [ ] Metrics instrumentation
- [ ] Distributed tracing
- [ ] Error tracking

## 🎉 Course Complete!

You've completed the entire Go Production Service Tutorial!

### What You've Mastered

✅ Go fundamentals (syntax through error handling)
✅ Intermediate Go (concurrency, testing, HTTP)
✅ Production APIs with Uber FX
✅ PostgreSQL and GORM
✅ Complete production service architecture
✅ Go best practices and standards
✅ Monorepo patterns
✅ Docker, performance, production patterns

### Next Steps in Your Go Journey

1. **Build a Real Project**: Apply what you've learned
2. **Contribute to Open Source**: golang/go, uber-go/fx, etc.
3. **Deep Dive**: Pick a topic and go deeper
4. **Stay Updated**: Follow Go blog and release notes

### Recommended Projects to Build

- RESTful API for a real use case
- Microservices with gRPC
- CLI tool with Cobra
- Real-time system with WebSockets
- Distributed system with message queues

## Resources for Continued Learning

- [Go Blog](https://go.dev/blog/)
- [Awesome Go](https://github.com/avelino/awesome-go)
- [Go Time Podcast](https://changelog.com/gotime)
- [GopherCon Talks](https://www.youtube.com/c/GopherAcademy)

**Congratulations on completing the course!** 🎊

## Further Reading

- [The Go Programming Language (book)](https://www.gopl.io/)
- [Learning Go (book)](https://www.oreilly.com/library/view/learning-go/9781492077206/)
- [Production-Ready Go](https://leanpub.com/production-ready-go)

