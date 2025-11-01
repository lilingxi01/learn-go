package main

import (
	"fmt"
	"math"
)

const (
	// Pi represents the mathematical constant π
	Pi = math.Pi
)

// CustomError demonstrates a custom error type with code and message
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError
func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Rectangle represents a rectangle shape with width and height
type Rectangle struct {
	Width, Height float64
}

// Area calculates and returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle shape with radius
type Circle struct {
	Radius float64
}

// Area calculates and returns the area of the circle
func (c Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

// Shape defines the interface for geometric shapes
type Shape interface {
	Area() float64
}

// Book represents a book with a title
type Book struct {
	Title string
}

func (b Book) String() string {
	return b.Title
}

func main() {
	fmt.Println("=== Type Assertions and Type Switches ===\n")

	// ===================================
	// 1. Basic Type Assertion
	// ===================================
	fmt.Println("1. Basic type assertion:")

	var i interface{} = "hello"

	// Type assertion (panics if wrong type)
	s := i.(string)
	fmt.Printf("   Asserted string: %s\n\n", s)

	// ===================================
	// 2. Safe Type Assertion
	// ===================================
	fmt.Println("2. Safe type assertion (with ok):")

	var value interface{} = 42

	// Safe assertion returns (value, bool)
	str, ok := value.(string)
	if ok {
		fmt.Printf("   Value is a string: %s\n", str)
	} else {
		fmt.Println("   Value is NOT a string")
	}

	num, ok := value.(int)
	if ok {
		fmt.Printf("   Value is an int: %d\n\n", num)
	}

	// ===================================
	// 3. Type Switch
	// ===================================
	fmt.Println("3. Type switch:")

	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"key": 100},
		nil,
	}

	for _, v := range values {
		describe(v)
	}
	fmt.Println()

	// ===================================
	// 4. Type Switch with Interface
	// ===================================
	fmt.Println("4. Type switch with custom interface:")

	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
	}

	for _, shape := range shapes {
		processShape(shape)
	}
	fmt.Println()

	// ===================================
	// 5. Empty Interface Array
	// ===================================
	fmt.Println("5. Working with []interface{}:")

	mixed := []interface{}{1, "two", 3.0, true}

	for i, v := range mixed {
		fmt.Printf("   Index %d: %v (type: %T)\n", i, v, v)
	}
	fmt.Println()

	// ===================================
	// 6. Type Assertion with Error Interface
	// ===================================
	fmt.Println("6. Type assertion with error interface:")

	var err error = CustomError{Code: 404, Message: "Not found"}

	// Type assert to get underlying type
	if ce, ok := err.(CustomError); ok {
		fmt.Printf("   Custom error code: %d\n", ce.Code)
		fmt.Printf("   Custom error message: %s\n\n", ce.Message)
	}

	// ===================================
	// 7. Interface Conversion
	// ===================================
	fmt.Println("7. Interface conversion:")

	var any interface{} = "hello world"

	// Convert to more specific interface
	var stringer fmt.Stringer = Book{Title: "Go Book"}

	fmt.Printf("   Stringer: %s\n", stringer)
	fmt.Printf("   Any: %v\n", any)

	fmt.Println("\n=== Type Assertions Tutorial Complete! ===")
}

// ===================================
// Helper Functions
// ===================================

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("   Integer: %d\n", v)
	case string:
		fmt.Printf("   String: %s\n", v)
	case float64:
		fmt.Printf("   Float: %.2f\n", v)
	case bool:
		fmt.Printf("   Boolean: %t\n", v)
	case []int:
		fmt.Printf("   Slice of ints: %v\n", v)
	case map[string]int:
		fmt.Printf("   Map: %v\n", v)
	case nil:
		fmt.Println("   Nil value")
	default:
		fmt.Printf("   Unknown type: %T\n", v)
	}
}

func processShape(s Shape) {
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("   Rectangle: %.0fx%.0f, Area: %.2f\n",
			shape.Width, shape.Height, shape.Area())
	case Circle:
		fmt.Printf("   Circle: radius %.0f, Area: %.2f\n",
			shape.Radius, shape.Area())
	default:
		fmt.Printf("   Unknown shape, Area: %.2f\n", s.Area())
	}
}
