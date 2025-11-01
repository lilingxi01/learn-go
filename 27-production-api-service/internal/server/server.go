package server

import (
	"context"
	"example.com/production-api/internal/config"
	"example.com/production-api/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// Server wraps the HTTP server
type Server struct {
	server *http.Server
}

// NewRouter creates the chi router with all routes
func NewRouter(userHandler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Production API Service"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.List)
			r.Post("/", userHandler.Create)
			r.Get("/{id}", userHandler.Get)
			r.Put("/{id}", userHandler.Update)
			r.Delete("/{id}", userHandler.Delete)
		})
	})

	return r
}

// New creates HTTP server with lifecycle
func New(lc fx.Lifecycle, cfg *config.Config, router chi.Router, logger zerolog.Logger) *Server {
	srv := &Server{
		server: &http.Server{
			Addr:    ":" + cfg.Server.Port,
			Handler: router,
		},
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info().
				Str("port", cfg.Server.Port).
				Msg("Starting HTTP server")

			go func() {
				if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error().Err(err).Msg("Server error")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info().Msg("Stopping HTTP server")
			return srv.server.Shutdown(ctx)
		},
	})

	return srv
}
