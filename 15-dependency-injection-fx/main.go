package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/fx"
)

// ===================================
// Logger Service
// ===================================

// Logger provides logging functionality
type Logger struct {
	prefix string
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		prefix: "[APP]",
	}
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	fmt.Printf("%s INFO: %s\n", l.prefix, msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	fmt.Printf("%s ERROR: %s\n", l.prefix, msg)
}

// ===================================
// HTTP Server
// ===================================

// Server wraps an HTTP server
type Server struct {
	server *http.Server
	logger *Logger
}

// NewServer creates a new HTTP server with injected dependencies
func NewServer(lc fx.Lifecycle, logger *Logger) *Server {
	logger.Info("Creating HTTP server")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("Request: %s %s", r.Method, r.URL.Path))
		fmt.Fprintf(w, "Hello from FX-powered server!\n")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK\n")
	})

	server := &Server{
		server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		logger: logger,
	}

	// Register lifecycle hooks
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server on :8080")
			go func() {
				if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error(fmt.Sprintf("Server error: %v", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server")
			return server.server.Shutdown(ctx)
		},
	})

	return server
}

// ===================================
// Database Service (Mock)
// ===================================

// Database represents a database connection
type Database struct {
	connectionString string
	logger           *Logger
}

// NewDatabase creates a new database connection with injected logger
func NewDatabase(lc fx.Lifecycle, logger *Logger) *Database {
	logger.Info("Creating database connection")

	db := &Database{
		connectionString: "postgresql://localhost:5432/mydb",
		logger:           logger,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Connecting to database: %s", db.connectionString))
			// In production: actually connect to database
			time.Sleep(100 * time.Millisecond) // Simulate connection
			logger.Info("Database connected")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing database connection")
			// In production: close database connection
			return nil
		},
	})

	return db
}

// Query simulates a database query
func (db *Database) Query(query string) {
	db.logger.Info(fmt.Sprintf("Executing query: %s", query))
}

// ===================================
// User Service
// ===================================

// UserService provides user-related operations
type UserService struct {
	db     *Database
	logger *Logger
}

// NewUserService creates a user service with injected dependencies
func NewUserService(db *Database, logger *Logger) *UserService {
	logger.Info("Creating UserService")
	return &UserService{
		db:     db,
		logger: logger,
	}
}

// GetUser retrieves a user (demonstration)
func (s *UserService) GetUser(id int) {
	s.logger.Info(fmt.Sprintf("Getting user %d", id))
	s.db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %d", id))
}

// ===================================
// Application Runner
// ===================================

// Run demonstrates using injected services
func Run(userService *UserService, logger *Logger) {
	logger.Info("Application started successfully")
	logger.Info("All dependencies injected and ready")

	// Demonstrate service usage
	userService.GetUser(1)

	logger.Info("Visit http://localhost:8080 to see the server")
	logger.Info("Press Ctrl+C to stop")
}

// ===================================
// Main - FX Application
// ===================================

func main() {
	app := fx.New(
		// Provide dependencies
		fx.Provide(
			NewLogger,
			NewDatabase,
			NewServer,
			NewUserService,
		),

		// Invoke to trigger initialization
		fx.Invoke(Run),
	)

	app.Run()
}
