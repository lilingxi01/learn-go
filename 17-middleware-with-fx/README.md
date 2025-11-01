# Middleware with FX

Master middleware patterns using Uber FX and chi router.

## Prerequisites

- Completed [16-api-with-fx-structure](../16-api-with-fx-structure/)
- Understanding of HTTP and FX

## What You'll Learn

- What middleware is and why it matters
- Creating middleware functions
- Middleware chains
- Authentication middleware
- Logging middleware with DI
- CORS handling
- Request ID propagation
- Recovery middleware

## What is Middleware?

Middleware is code that runs before your handler:

```
Request → Middleware 1 → Middleware 2 → Handler → Response
```

Common uses:

- Authentication/authorization
- Logging
- CORS
- Rate limiting
- Request ID tracking
- Panic recovery

## Middleware Signature

```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before handler
        next.ServeHTTP(w, r)
        // After handler
    })
}
```

## Running the Example

```bash
cd 17-middleware-with-fx
go mod download
go run .
```

## Next Steps

Move to **18-request-validation** for input validation patterns

## Further Reading

- [Writing Middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
- [Chi Middleware](https://github.com/go-chi/chi#middlewares)
