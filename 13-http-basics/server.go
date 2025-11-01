package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Carol", Email: "carol@example.com"},
}

func main() {
	fmt.Println("=== Go HTTP Server Tutorial ===\n")

	// ===================================
	// Register Handlers
	// ===================================
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/user/", userHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/headers", headersHandler)

	// ===================================
	// Start Server
	// ===================================
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  http://localhost:8080/")
	fmt.Println("  GET  http://localhost:8080/about")
	fmt.Println("  GET  http://localhost:8080/api/users")
	fmt.Println("  POST http://localhost:8080/echo")
	fmt.Println("  GET  http://localhost:8080/headers")
	fmt.Println("\nPress Ctrl+C to stop the server")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// homeHandler handles requests to the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Welcome to Go HTTP Server!\n")
	fmt.Fprintf(w, "Visit /about or /api/users for more\n")
}

// aboutHandler provides information about the server
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"name":    "Go Tutorial Server",
		"version": "1.0.0",
		"uptime":  time.Since(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// usersHandler handles user list requests
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return list of users
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		// Create new user
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newUser.ID = len(users) + 1
		users = append(users, newUser)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// userHandler handles individual user requests
func userHandler(w http.ResponseWriter, r *http.Request) {
	// In production, use a proper router for path parameters
	// This is a simple demonstration
	fmt.Fprintf(w, "User detail endpoint\n")
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
}

// echoHandler echoes back the request body
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and echo body
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"echo":      data,
		"timestamp": time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// headersHandler displays request headers
func headersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	headers := make(map[string][]string)
	for name, values := range r.Header {
		headers[name] = values
	}

	response := map[string]interface{}{
		"method":  r.Method,
		"path":    r.URL.Path,
		"headers": headers,
	}

	json.NewEncoder(w).Encode(response)
}

var startTime = time.Now()
