package main

import (
	"fmt"
	"math"
)

const (
	// Pi is a more precise value of π for calculations
	Pi = math.Pi
)

// Shape defines the interface for geometric shapes.
// Any type that implements Area, Perimeter, and Name satisfies this interface.
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Rectangle type
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

// Circle type
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Name() string {
	return "Circle"
}

// Triangle type
type Triangle struct {
	Base, Height, SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t Triangle) Name() string {
	return "Triangle"
}

// ===================================
// Functions using interfaces
// ===================================

func printShapeInfo(s Shape) {
	fmt.Printf("%s: Area=%.2f, Perimeter=%.2f\n",
		s.Name(), s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// ===================================
// Other interfaces
// ===================================

// Stringer interface (from fmt package)
type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("'%s' by %s (%d pages)", b.Title, b.Author, b.Pages)
}

// Custom interface for printing
type Printer interface {
	Print()
}

type Document struct {
	Content string
}

func (d Document) Print() {
	fmt.Println("Document:", d.Content)
}

type Photo struct {
	Filename string
}

func (p Photo) Print() {
	fmt.Println("Photo:", p.Filename)
}

func printAll(items []Printer) {
	for _, item := range items {
		item.Print()
	}
}

// File represents a file with content
type File struct {
	content string
}

// Read returns the file content
func (f *File) Read() string {
	return f.content
}

// Write updates the file content
func (f *File) Write(content string) {
	f.content = content
}

func main() {
	fmt.Println("=== Go Interfaces Tutorial ===\n")

	// ===================================
	// 1. Basic Interface Usage
	// ===================================
	fmt.Println("1. Basic interface usage:")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{Base: 6, Height: 8, SideA: 6, SideB: 8, SideC: 10}

	printShapeInfo(rect)
	printShapeInfo(circle)
	printShapeInfo(triangle)
	fmt.Println()

	// ===================================
	// 2. Slice of Interfaces
	// ===================================
	fmt.Println("2. Slice of interfaces:")

	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Triangle{Base: 6, Height: 4, SideA: 5, SideB: 5, SideC: 6},
	}

	total := totalArea(shapes)
	fmt.Printf("   Total area of all shapes: %.2f\n\n", total)

	// ===================================
	// 3. Empty Interface (interface{})
	// ===================================
	fmt.Println("3. Empty interface:")

	var i interface{}

	i = 42
	fmt.Printf("   i = %v (type: %T)\n", i, i)

	i = "hello"
	fmt.Printf("   i = %v (type: %T)\n", i, i)

	i = []int{1, 2, 3}
	fmt.Printf("   i = %v (type: %T)\n\n", i, i)

	// ===================================
	// 4. Implementing fmt.Stringer
	// ===================================
	fmt.Println("4. Implementing fmt.Stringer:")

	book := Book{
		Title:  "The Go Programming Language",
		Author: "Donovan & Kernighan",
		Pages:  400,
	}

	// String() method called automatically by fmt
	fmt.Printf("   Book: %s\n\n", book)

	// ===================================
	// 5. Multiple Types, Same Interface
	// ===================================
	fmt.Println("5. Multiple types implementing same interface:")

	items := []Printer{
		Document{Content: "Important memo"},
		Photo{Filename: "vacation.jpg"},
		Document{Content: "Meeting notes"},
	}

	printAll(items)
	fmt.Println()

	// ===================================
	// 6. Interface Composition
	// ===================================
	fmt.Println("6. Interface composition:")

	// File implements ReadWriter interface
	file := &File{content: "initial content"}

	fmt.Printf("   Read: %s\n", file.Read())
	file.Write("updated content")
	fmt.Printf("   After write: %s\n\n", file.Read())

	// ===================================
	// 7. Nil Interface
	// ===================================
	fmt.Println("7. Nil interface:")

	var s Shape
	fmt.Printf("   Shape is nil: %t\n", s == nil)

	if s == nil {
		fmt.Println("   Cannot call methods on nil interface")
	}
	fmt.Println()

	// ===================================
	// 8. Interface with Value and Pointer Receivers
	// ===================================
	fmt.Println("8. Value vs pointer receivers:")

	rectValue := Rectangle{Width: 10, Height: 5}
	rectPointer := &Rectangle{Width: 10, Height: 5}

	// Both work because Go automatically handles it
	printShapeInfo(rectValue)
	printShapeInfo(rectPointer)

	fmt.Println("\n=== Interfaces Tutorial Complete! ===")
}
