package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// In-memory storage (in production, use a database)
var (
	users  = make(map[int]User)
	nextID = 1
)

func main() {
	// Initialize with sample data
	users[nextID] = User{
		ID:        nextID,
		Name:      "Alice Johnson",
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
	}
	nextID++

	users[nextID] = User{
		ID:        nextID,
		Name:      "Bob Smith",
		Email:     "bob@example.com",
		CreatedAt: time.Now(),
	}
	nextID++

	// Setup router
	r := setupRouter()

	// Start server
	port := ":8080"
	fmt.Printf("🚀 REST API Server starting on http://localhost%s\n\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    http://localhost:8080/api/users       - List all users")
	fmt.Println("  POST   http://localhost:8080/api/users       - Create user")
	fmt.Println("  GET    http://localhost:8080/api/users/{id}  - Get user by ID")
	fmt.Println("  PUT    http://localhost:8080/api/users/{id}  - Update user")
	fmt.Println("  DELETE http://localhost:8080/api/users/{id}  - Delete user")
	fmt.Println("\nPress Ctrl+C to stop")

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

// setupRouter configures and returns the chi router
func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Routes
	r.Get("/", homeHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", listUsers)
			r.Post("/", createUser)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", getUser)
				r.Put("/", updateUser)
				r.Delete("/", deleteUser)
			})
		})
	})

	return r
}

// homeHandler responds with API information
func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"name":    "Go REST API Tutorial",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"users": "/api/users",
		},
	}

	respondJSON(w, http.StatusOK, response)
}

// listUsers returns all users
func listUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	respondJSON(w, http.StatusOK, userList)
}

// createUser creates a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Validate input
	if user.Name == "" {
		respondError(w, http.StatusBadRequest, "Validation failed", "name is required")
		return
	}
	if user.Email == "" {
		respondError(w, http.StatusBadRequest, "Validation failed", "email is required")
		return
	}

	// Create user
	user.ID = nextID
	user.CreatedAt = time.Now()
	users[nextID] = user
	nextID++

	respondJSON(w, http.StatusCreated, user)
}

// getUser returns a single user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	user, exists := users[id]
	if !exists {
		respondError(w, http.StatusNotFound, "User not found", fmt.Sprintf("user with ID %d does not exist", id))
		return
	}

	respondJSON(w, http.StatusOK, user)
}

// updateUser updates an existing user
func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	user, exists := users[id]
	if !exists {
		respondError(w, http.StatusNotFound, "User not found", fmt.Sprintf("user with ID %d does not exist", id))
		return
	}

	var updates User
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Update fields (keeping ID and CreatedAt)
	user.Name = updates.Name
	user.Email = updates.Email
	users[id] = user

	respondJSON(w, http.StatusOK, user)
}

// deleteUser deletes a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	if _, exists := users[id]; !exists {
		respondError(w, http.StatusNotFound, "User not found", fmt.Sprintf("user with ID %d does not exist", id))
		return
	}

	delete(users, id)

	w.WriteHeader(http.StatusNoContent)
}

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// respondError sends an error response
func respondError(w http.ResponseWriter, status int, error string, message string) {
	respondJSON(w, status, ErrorResponse{
		Error:   error,
		Message: message,
	})
}
