package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== Production Patterns ===\n")

	// Create HTTP server
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", healthHandler)

	// Readiness check endpoint
	mux.HandleFunc("/ready", readinessHandler)

	// API endpoint
	mux.HandleFunc("/api/data", dataHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		fmt.Println("🚀 Server starting on :8080")
		fmt.Println("   GET /health - Health check")
		fmt.Println("   GET /ready  - Readiness check")
		fmt.Println("   GET /api/data - API endpoint")
		fmt.Println("\nPress Ctrl+C for graceful shutdown")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("✓ Server stopped gracefully")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy"}`)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Check dependencies (database, cache, etc.)
	// For demo, always ready
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ready"}`)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Add timeout to request
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Simulate work
	select {
	case <-time.After(100 * time.Millisecond):
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"data":"response"}`)
	case <-ctx.Done():
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
	}
}
