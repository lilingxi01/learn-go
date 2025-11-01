package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Config holds application configuration
type Config struct {
	Port string
}

// NewConfig creates configuration
func NewConfig() *Config {
	return &Config{Port: ":8080"}
}

// NewLogger creates a zap logger
func NewLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// NewRouter creates chi router with middleware
func NewRouter(logger *zap.Logger) chi.Router {
	r := chi.NewRouter()

	// Built-in middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// Custom middleware with injected logger
	r.Use(LoggingMiddleware(logger))

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Middleware demonstration with FX",
		})
	})

	r.Get("/api/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected endpoint"))
	})

	return r
}

// LoggingMiddleware creates a logging middleware with injected logger
func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			logger.Info("request completed",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}

// NewServer creates HTTP server with lifecycle
func NewServer(lc fx.Lifecycle, cfg *Config, router chi.Router, logger *zap.Logger) *http.Server {
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting server", zap.String("port", cfg.Port))
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping server")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			NewRouter,
			NewServer,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
