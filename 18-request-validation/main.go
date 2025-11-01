package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

// User represents a user with validation tags
type User struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=0,lte=120"`
}

// Config holds configuration
type Config struct {
	Port string
}

// NewConfig provides configuration
func NewConfig() *Config {
	return &Config{Port: ":8080"}
}

// NewValidator provides a validator instance
func NewValidator() *validator.Validate {
	return validator.New()
}

// Handler contains validation logic
type Handler struct {
	validate *validator.Validate
}

// NewHandler creates handler with injected validator
func NewHandler(v *validator.Validate) *Handler {
	return &Handler{validate: v}
}

// CreateUser handles user creation with validation
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON", nil)
		return
	}

	// Validate
	if err := h.validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		respondError(w, http.StatusBadRequest, "validation failed", formatValidationErrors(validationErrors))
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

// NewRouter creates router with handlers
func NewRouter(handler *Handler) chi.Router {
	r := chi.NewRouter()
	r.Post("/api/users", handler.CreateUser)
	return r
}

// NewServer creates HTTP server
func NewServer(lc fx.Lifecycle, cfg *Config, router chi.Router) *http.Server {
	srv := &http.Server{Addr: cfg.Port, Handler: router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("Server starting on %s\n", cfg.Port)
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewValidator,
			NewHandler,
			NewRouter,
			NewServer,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string, details interface{}) {
	respondJSON(w, status, map[string]interface{}{
		"error":   message,
		"details": details,
	})
}

func formatValidationErrors(errs validator.ValidationErrors) []map[string]string {
	var errors []map[string]string
	for _, err := range errs {
		errors = append(errors, map[string]string{
			"field": err.Field(),
			"tag":   err.Tag(),
			"value": fmt.Sprintf("%v", err.Value()),
		})
	}
	return errors
}
