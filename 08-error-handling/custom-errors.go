package main

import (
	"errors"
	"fmt"
)

// ===================================
// Custom Error Types
// ===================================

// NetworkError represents a network-related error
type NetworkError struct {
	Code    int
	Message string
	URL     string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error %d: %s (URL: %s)", e.Code, e.Message, e.URL)
}

// Temporary method makes it compatible with temporary error checking
func (e *NetworkError) Temporary() bool {
	return e.Code >= 500
}

// ValidationError for input validation
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s (value: %v)",
		e.Field, e.Message, e.Value)
}

// DatabaseError for database operations
type DatabaseError struct {
	Operation string
	Table     string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s on %s: %v",
		e.Operation, e.Table, e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// MultiError aggregates multiple errors
type MultiError []error

func (m MultiError) Error() string {
	if len(m) == 0 {
		return "no errors"
	}
	if len(m) == 1 {
		return m[0].Error()
	}

	msg := fmt.Sprintf("%d errors occurred:", len(m))
	for i, err := range m {
		msg += fmt.Sprintf("\n  %d) %v", i+1, err)
	}
	return msg
}

func main() {
	fmt.Println("=== Custom Error Types Tutorial ===\n")

	// ===================================
	// 1. NetworkError
	// ===================================
	fmt.Println("1. Network errors:")

	err := fetchData("https://api.example.com/users")
	if err != nil {
		if ne, ok := err.(*NetworkError); ok {
			fmt.Printf("   Network error occurred\n")
			fmt.Printf("   Code: %d\n", ne.Code)
			fmt.Printf("   Message: %s\n", ne.Message)
			fmt.Printf("   URL: %s\n", ne.URL)
			fmt.Printf("   Temporary: %t\n", ne.Temporary())
		}
	}
	fmt.Println()

	// ===================================
	// 2. ValidationError
	// ===================================
	fmt.Println("2. Validation errors:")

	user := User{Name: "", Age: -5, Email: "invalid"}
	if err := validateUser(user); err != nil {
		fmt.Printf("   Validation failed:\n")

		// Check if it's a validation error
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("   Field: %s\n", ve.Field)
			fmt.Printf("   Message: %s\n", ve.Message)
		} else {
			fmt.Printf("   %v\n", err)
		}
	}
	fmt.Println()

	// ===================================
	// 3. DatabaseError with Wrapping
	// ===================================
	fmt.Println("3. Database errors:")

	err = deleteUser(999)
	if err != nil {
		fmt.Printf("   %v\n", err)

		// Check underlying error
		var dbe *DatabaseError
		if errors.As(err, &dbe) {
			fmt.Printf("   Operation: %s\n", dbe.Operation)
			fmt.Printf("   Table: %s\n", dbe.Table)

			// Check if wrapped error is not found
			if errors.Is(dbe.Err, ErrNotFound) {
				fmt.Println("   Underlying cause: Record not found")
			}
		}
	}
	fmt.Println()

	// ===================================
	// 4. MultiError
	// ===================================
	fmt.Println("4. Multiple errors:")

	errs := processMultiple([]int{1, 0, 3, 0, 5})
	if len(errs) > 0 {
		fmt.Printf("   %v\n", MultiError(errs))
	}
	fmt.Println()

	// ===================================
	// 5. Error Type Checking
	// ===================================
	fmt.Println("5. Error type checking patterns:")

	errors := []error{
		&NetworkError{Code: 500, Message: "Server error", URL: "api.com"},
		&ValidationError{Field: "email", Message: "invalid format"},
		&DatabaseError{Operation: "INSERT", Table: "users", Err: errors.New("duplicate")},
	}

	for i, err := range errors {
		fmt.Printf("   Error %d: ", i+1)

		switch e := err.(type) {
		case *NetworkError:
			fmt.Printf("Network (Code: %d)\n", e.Code)
		case *ValidationError:
			fmt.Printf("Validation (Field: %s)\n", e.Field)
		case *DatabaseError:
			fmt.Printf("Database (Operation: %s)\n", e.Operation)
		default:
			fmt.Printf("Unknown\n")
		}
	}
	fmt.Println()

	// ===================================
	// 6. Error with Context
	// ===================================
	fmt.Println("6. Error with context:")

	err = processRequest("api.example.com", "/users/1")
	if err != nil {
		fmt.Printf("   %v\n", err)
	}

	fmt.Println("\n=== Custom Errors Tutorial Complete! ===")
}

// ===================================
// Helper Functions and Types
// ===================================

func fetchData(url string) error {
	// Simulate network error
	return &NetworkError{
		Code:    503,
		Message: "Service Unavailable",
		URL:     url,
	}
}

type User struct {
	Name  string
	Age   int
	Email string
}

func validateUser(u User) error {
	if u.Name == "" {
		return &ValidationError{
			Field:   "name",
			Value:   u.Name,
			Message: "name cannot be empty",
		}
	}
	if u.Age < 0 {
		return &ValidationError{
			Field:   "age",
			Value:   u.Age,
			Message: "age cannot be negative",
		}
	}
	return nil
}

var ErrNotFound = errors.New("record not found")

func deleteUser(id int) error {
	// Simulate not found
	if id == 999 {
		return &DatabaseError{
			Operation: "DELETE",
			Table:     "users",
			Err:       ErrNotFound,
		}
	}
	return nil
}

func processMultiple(numbers []int) []error {
	var errs []error

	for i, num := range numbers {
		if num == 0 {
			errs = append(errs, fmt.Errorf("invalid value at index %d", i))
		}
	}

	return errs
}

func processRequest(host, path string) error {
	// Simulate layered errors
	err := &NetworkError{Code: 404, Message: "Not Found", URL: host + path}
	return fmt.Errorf("failed to process request to %s: %w", host, err)
}
