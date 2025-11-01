// Package server provides the HTTP server with lifecycle management
package server

import (
	"context"
	"example.com/api-with-fx/config"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

// Server wraps the HTTP server
type Server struct {
	server *http.Server
	config *Config
}

// New creates a new HTTP server with injected dependencies and lifecycle hooks
func New(lc fx.Lifecycle, cfg *config.Config, router chi.Router) *Server {
	fmt.Println("🚀 Creating HTTP server")

	srv := &Server{
		server: &http.Server{
			Addr:    cfg.ServerPort,
			Handler: router,
		},
		config: cfg,
	}

	// Register lifecycle hooks
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("▶️  Starting %s v%s on %s\n", cfg.AppName, cfg.Version, cfg.ServerPort)

			// Start server in goroutine
			go func() {
				if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Printf("❌ Server error: %v\n", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("⏹  Stopping HTTP server")
			return srv.server.Shutdown(ctx)
		},
	})

	return srv
}
