package handlers

import (
	"encoding/json"
	"example.com/production-api/internal/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	db       *gorm.DB
	validate *validator.Validate
}

// NewUserHandler creates a new user handler with injected dependencies
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db:       db,
		validate: validator.New(),
	}
}

// List returns all users
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := h.db.Find(&users)
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "database error")
		return
	}

	respondJSON(w, http.StatusOK, users)
}

// Get returns a single user
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	var user models.User
	result := h.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			respondError(w, http.StatusNotFound, "user not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "database error")
		return
	}

	respondJSON(w, http.StatusOK, user)
}

// Create creates a new user
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if err := h.validate.Struct(user); err != nil {
		respondError(w, http.StatusBadRequest, "validation failed")
		return
	}

	result := h.db.Create(&user)
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

// Update updates an existing user
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	var user models.User
	if result := h.db.First(&user, id); result.Error != nil {
		respondError(w, http.StatusNotFound, "user not found")
		return
	}

	var updates models.User
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	h.db.Model(&user).Updates(updates)
	respondJSON(w, http.StatusOK, user)
}

// Delete deletes a user
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	result := h.db.Delete(&models.User{}, id)
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "failed to delete")
		return
	}

	if result.RowsAffected == 0 {
		respondError(w, http.StatusNotFound, "user not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
