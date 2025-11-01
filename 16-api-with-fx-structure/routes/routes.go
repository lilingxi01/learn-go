// Package routes provides HTTP route registration
package routes

import (
	"example.com/api-with-fx/handlers"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New creates and configures the chi router with all routes
func New(userHandler *handlers.UserHandler) chi.Router {
	fmt.Println("🛣  Registering routes")

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Root endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API powered by Uber FX"))
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.List)
			r.Post("/", userHandler.Create)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", userHandler.Get)
				r.Put("/", userHandler.Update)
				r.Delete("/", userHandler.Delete)
			})
		})
	})

	return r
}
