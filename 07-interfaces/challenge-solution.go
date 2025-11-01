package main

import (
	"fmt"
	"math"
)

// ===================================
// Interfaces
// ===================================

// Shape interface defines geometric shape behavior
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Drawable interface for shapes that can be drawn
type Drawable interface {
	Draw()
}

// Comparable interface for comparing shapes
type Comparable interface {
	IsLargerThan(other Shape) bool
}

// ===================================
// Rectangle Implementation
// ===================================

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Draw() {
	fmt.Println("   Drawing Rectangle...")
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1fx%.1f)", r.Width, r.Height)
}

// ===================================
// Circle Implementation
// ===================================

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Draw() {
	fmt.Println("   Drawing Circle...")
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.1f)", c.Radius)
}

// ===================================
// Triangle Implementation
// ===================================

type Triangle struct {
	Base, Height        float64 // For area calculation
	SideA, SideB, SideC float64 // For perimeter
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t Triangle) Draw() {
	fmt.Println("   Drawing Triangle...")
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(base=%.1f, height=%.1f)", t.Base, t.Height)
}

// ===================================
// Shape Collection (Bonus)
// ===================================

type ShapeCollection struct {
	shapes []Shape
}

func NewShapeCollection() *ShapeCollection {
	return &ShapeCollection{shapes: []Shape{}}
}

func (sc *ShapeCollection) Add(s Shape) {
	sc.shapes = append(sc.shapes, s)
}

func (sc *ShapeCollection) TotalArea() float64 {
	total := 0.0
	for _, shape := range sc.shapes {
		total += shape.Area()
	}
	return total
}

func (sc *ShapeCollection) PrintAll() {
	for i, shape := range sc.shapes {
		fmt.Printf("   %d. %v - Area: %.2f, Perimeter: %.2f\n",
			i+1, shape, shape.Area(), shape.Perimeter())
	}
}

// ===================================
// Required Functions
// ===================================

func CalculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func FindLargestShape(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}

	largest := shapes[0]
	largestArea := largest.Area()

	for _, shape := range shapes[1:] {
		area := shape.Area()
		if area > largestArea {
			largest = shape
			largestArea = area
		}
	}

	return largest
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("   Shape: %v\n", s)
	fmt.Printf("   Area: %.2f\n", s.Area())
	fmt.Printf("   Perimeter: %.2f\n", s.Perimeter())

	// Type-specific details using type assertion
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("   Type: Rectangle (%.1f x %.1f)\n", shape.Width, shape.Height)
	case Circle:
		fmt.Printf("   Type: Circle (radius: %.1f)\n", shape.Radius)
		fmt.Printf("   Diameter: %.2f\n", 2*shape.Radius)
	case Triangle:
		fmt.Printf("   Type: Triangle (base: %.1f, height: %.1f)\n", shape.Base, shape.Height)
	}

	// Check if drawable
	if drawable, ok := s.(Drawable); ok {
		drawable.Draw()
	}
}

func main() {
	fmt.Println("=== Shape Calculator Challenge Solution ===\n")

	// ===================================
	// 1. Create Shapes
	// ===================================
	fmt.Println("1. Creating shapes:")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{
		Base: 6, Height: 8,
		SideA: 6, SideB: 8, SideC: 10,
	}

	shapes := []Shape{rect, circle, triangle}

	for _, shape := range shapes {
		fmt.Printf("   Created: %v\n", shape)
	}
	fmt.Println()

	// ===================================
	// 2. Calculate Total Area
	// ===================================
	fmt.Println("2. Calculate total area:")
	totalArea := CalculateTotalArea(shapes)
	fmt.Printf("   Total area of all shapes: %.2f\n\n", totalArea)

	// ===================================
	// 3. Find Largest Shape
	// ===================================
	fmt.Println("3. Find largest shape:")
	largest := FindLargestShape(shapes)
	fmt.Printf("   Largest shape: %v\n", largest)
	fmt.Printf("   Area: %.2f\n\n", largest.Area())

	// ===================================
	// 4. Print Details for Each Shape
	// ===================================
	fmt.Println("4. Detailed information for each shape:")
	for i, shape := range shapes {
		fmt.Printf("\nShape %d:\n", i+1)
		PrintShapeDetails(shape)
	}
	fmt.Println()

	// ===================================
	// Bonus: Drawable Interface
	// ===================================
	fmt.Println("5. Drawing all shapes:")
	for _, shape := range shapes {
		if drawable, ok := shape.(Drawable); ok {
			drawable.Draw()
		}
	}
	fmt.Println()

	// ===================================
	// Bonus: Shape Collection
	// ===================================
	fmt.Println("6. Using ShapeCollection:")
	collection := NewShapeCollection()

	collection.Add(Rectangle{Width: 5, Height: 3})
	collection.Add(Circle{Radius: 4})
	collection.Add(Triangle{Base: 6, Height: 4, SideA: 5, SideB: 5, SideC: 6})
	collection.Add(Rectangle{Width: 8, Height: 4})

	collection.PrintAll()
	fmt.Printf("\n   Total area in collection: %.2f\n\n", collection.TotalArea())

	// ===================================
	// Bonus: Type Switch Example
	// ===================================
	fmt.Println("7. Type-specific operations:")

	for _, shape := range shapes {
		switch s := shape.(type) {
		case Rectangle:
			diagonal := math.Sqrt(s.Width*s.Width + s.Height*s.Height)
			fmt.Printf("   Rectangle diagonal: %.2f\n", diagonal)
		case Circle:
			circumference := s.Perimeter()
			fmt.Printf("   Circle circumference: %.2f\n", circumference)
		case Triangle:
			fmt.Printf("   Triangle perimeter: %.2f\n", s.Perimeter())
		}
	}

	// ===================================
	// Summary
	// ===================================
	fmt.Println("\n=== What You Learned ===")
	fmt.Println("✓ Defining and implementing interfaces")
	fmt.Println("✓ Implicit interface implementation")
	fmt.Println("✓ Using interfaces as function parameters")
	fmt.Println("✓ Type assertions and type switches")
	fmt.Println("✓ Multiple interfaces for same type")
	fmt.Println("✓ Implementing fmt.Stringer")
	fmt.Println("✓ Interface composition")
	fmt.Println("✓ Real-world interface patterns")
}
