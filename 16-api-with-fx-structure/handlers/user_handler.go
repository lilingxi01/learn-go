// Package handlers provides HTTP request handlers
package handlers

import (
	"encoding/json"
	"example.com/api-with-fx/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// User represents a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserHandler handles user-related requests
type UserHandler struct {
	config *config.Config
	users  map[int]User
	nextID int
}

// NewUserHandler creates a new user handler with injected configuration
func NewUserHandler(cfg *config.Config) *UserHandler {
	fmt.Println("👥 Creating UserHandler")

	handler := &UserHandler{
		config: cfg,
		users:  make(map[int]User),
		nextID: 1,
	}

	// Initialize with sample data
	handler.users[handler.nextID] = User{
		ID:    handler.nextID,
		Name:  "Alice",
		Email: "alice@example.com",
	}
	handler.nextID++

	return handler
}

// List returns all users
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0, len(h.users))
	for _, user := range h.users {
		userList = append(userList, user)
	}

	h.respondJSON(w, http.StatusOK, userList)
}

// Get returns a single user
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	user, exists := h.users[id]
	if !exists {
		h.respondError(w, http.StatusNotFound, "user not found")
		return
	}

	h.respondJSON(w, http.StatusOK, user)
}

// Create creates a new user
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	// Validate
	if user.Name == "" || user.Email == "" {
		h.respondError(w, http.StatusBadRequest, "name and email are required")
		return
	}

	// Create user
	user.ID = h.nextID
	h.users[h.nextID] = user
	h.nextID++

	h.respondJSON(w, http.StatusCreated, user)
}

// Update updates an existing user
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	user, exists := h.users[id]
	if !exists {
		h.respondError(w, http.StatusNotFound, "user not found")
		return
	}

	var updates User
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	user.Name = updates.Name
	user.Email = updates.Email
	h.users[id] = user

	h.respondJSON(w, http.StatusOK, user)
}

// Delete deletes a user
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	if _, exists := h.users[id]; !exists {
		h.respondError(w, http.StatusNotFound, "user not found")
		return
	}

	delete(h.users, id)
	w.WriteHeader(http.StatusNoContent)
}

// Helper methods

func (h *UserHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}
