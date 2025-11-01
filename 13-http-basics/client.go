package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// User represents a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("=== Go HTTP Client Tutorial ===\n")

	// Base URL (start server.go first!)
	baseURL := "http://localhost:8080"

	// ===================================
	// 1. Simple GET Request
	// ===================================
	fmt.Println("1. Simple GET request:")

	resp, err := http.Get(baseURL + "/")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		fmt.Println("   (Make sure server.go is running!)\n")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("   Error reading body: %v\n", err)
		return
	}

	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Printf("   Body: %s\n", string(body))

	// ===================================
	// 2. GET with JSON Response
	// ===================================
	fmt.Println("2. GET request with JSON parsing:")

	resp, err = http.Get(baseURL + "/api/users")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Printf("   Error decoding JSON: %v\n", err)
		return
	}

	fmt.Printf("   Received %d users:\n", len(users))
	for _, user := range users {
		fmt.Printf("   - %s (ID: %d)\n", user.Name, user.ID)
	}
	fmt.Println()

	// ===================================
	// 3. POST Request with JSON
	// ===================================
	fmt.Println("3. POST request with JSON body:")

	newUser := User{
		Name:  "David",
		Email: "david@example.com",
	}

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		fmt.Printf("   Error marshaling JSON: %v\n", err)
		return
	}

	resp, err = http.Post(
		baseURL+"/api/users",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var createdUser User
	if err := json.NewDecoder(resp.Body).Decode(&createdUser); err != nil {
		fmt.Printf("   Error decoding response: %v\n", err)
		return
	}

	fmt.Printf("   Created user: %s (ID: %d)\n\n", createdUser.Name, createdUser.ID)

	// ===================================
	// 4. Custom Request with Headers
	// ===================================
	fmt.Println("4. Custom request with headers:")

	req, err := http.NewRequest("GET", baseURL+"/api/users", nil)
	if err != nil {
		fmt.Printf("   Error creating request: %v\n", err)
		return
	}

	req.Header.Set("User-Agent", "Go-Tutorial-Client/1.0")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Printf("   Content-Type: %s\n\n", resp.Header.Get("Content-Type"))

	// ===================================
	// 5. Query Parameters
	// ===================================
	fmt.Println("5. Request with query parameters:")

	req2, err := http.NewRequest("GET", baseURL+"/search", nil)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}

	// Add query parameters
	q := req2.URL.Query()
	q.Add("q", "golang")
	q.Add("limit", "10")
	req2.URL.RawQuery = q.Encode()

	fmt.Printf("   URL: %s\n\n", req2.URL.String())

	// ===================================
	// 6. Handling Different Status Codes
	// ===================================
	fmt.Println("6. Handling HTTP status codes:")

	resp, err = http.Get(baseURL + "/nonexistent")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("   Success!")
	case http.StatusNotFound:
		fmt.Println("   Page not found (404)")
	case http.StatusInternalServerError:
		fmt.Println("   Server error (500)")
	default:
		fmt.Printf("   Status: %d\n", resp.StatusCode)
	}
	fmt.Println()

	// ===================================
	// 7. Custom HTTP Client
	// ===================================
	fmt.Println("7. Custom HTTP client with timeout:")

	customClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: false,
		},
	}

	req, _ = http.NewRequest("GET", baseURL+"/about", nil)
	resp, err = customClient.Do(req)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var about map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&about)
	fmt.Printf("   Server info: %v\n", about)

	fmt.Println("\n=== HTTP Client Tutorial Complete! ===")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("✓ Use http.Get for simple GET requests")
	fmt.Println("✓ Use http.Post for POST requests")
	fmt.Println("✓ Use http.NewRequest for custom requests")
	fmt.Println("✓ Always defer resp.Body.Close()")
	fmt.Println("✓ Set timeouts to avoid hanging")
	fmt.Println("✓ Check status codes before processing")
	fmt.Println("✓ Use json.Decoder for JSON responses")
}
