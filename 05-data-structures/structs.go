package main

import "fmt"

// Define structs at package level
type Person struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Person   Person  // Embedded struct
	Address  Address // Embedded struct
	JobTitle string
	Salary   float64
}

// Method with value receiver (read-only)
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Method with value receiver
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Method with pointer receiver (can modify struct)
func (p *Person) HaveBirthday() {
	p.Age++
}

// Method with pointer receiver
func (p *Person) UpdateEmail(email string) {
	p.Email = email
}

// Method on Employee
func (e Employee) GetFullInfo() string {
	return fmt.Sprintf("%s (%d) - %s, works as %s",
		e.Person.Name, e.Person.Age, e.Address.City, e.JobTitle)
}

func main() {
	fmt.Println("=== Go Structs Tutorial ===\n")

	// ===================================
	// 1. Creating Structs
	// ===================================
	fmt.Println("1. Creating structs:")

	// Using field names (preferred)
	alice := Person{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}
	fmt.Printf("   %+v\n", alice)

	// Positional (not recommended)
	bob := Person{"Bob", 30, "bob@example.com"}
	fmt.Printf("   %+v\n\n", bob)

	// ===================================
	// 2. Accessing Fields
	// ===================================
	fmt.Println("2. Accessing struct fields:")
	fmt.Printf("   Name: %s\n", alice.Name)
	fmt.Printf("   Age: %d\n", alice.Age)
	fmt.Printf("   Email: %s\n\n", alice.Email)

	// ===================================
	// 3. Modifying Fields
	// ===================================
	fmt.Println("3. Modifying struct fields:")
	carol := Person{Name: "Carol", Age: 28, Email: "carol@example.com"}
	fmt.Printf("   Before: %+v\n", carol)

	carol.Age = 29
	carol.Email = "carol.new@example.com"
	fmt.Printf("   After: %+v\n\n", carol)

	// ===================================
	// 4. Struct Pointers
	// ===================================
	fmt.Println("4. Struct pointers:")
	dave := &Person{Name: "Dave", Age: 35, Email: "dave@example.com"}
	fmt.Printf("   Pointer: %p\n", dave)
	fmt.Printf("   Value: %+v\n", *dave)
	fmt.Printf("   Field access: %s (auto-dereferenced)\n\n", dave.Name)

	// ===================================
	// 5. Zero Value Struct
	// ===================================
	fmt.Println("5. Zero value struct:")
	var empty Person
	fmt.Printf("   Empty person: %+v\n", empty)
	fmt.Printf("   Name: '%s', Age: %d, Email: '%s'\n\n",
		empty.Name, empty.Age, empty.Email)

	// ===================================
	// 6. Methods (Value Receiver)
	// ===================================
	fmt.Println("6. Methods with value receiver:")
	fmt.Printf("   %s\n", alice.Greet())
	fmt.Printf("   Is Alice an adult? %t\n\n", alice.IsAdult())

	// ===================================
	// 7. Methods (Pointer Receiver)
	// ===================================
	fmt.Println("7. Methods with pointer receiver:")
	eve := Person{Name: "Eve", Age: 24, Email: "eve@example.com"}
	fmt.Printf("   Before birthday: Age = %d\n", eve.Age)

	eve.HaveBirthday()
	fmt.Printf("   After birthday: Age = %d\n", eve.Age)

	eve.UpdateEmail("eve.new@example.com")
	fmt.Printf("   Updated email: %s\n\n", eve.Email)

	// ===================================
	// 8. Nested Structs
	// ===================================
	fmt.Println("8. Nested structs:")

	address := Address{
		Street:  "123 Main St",
		City:    "San Francisco",
		ZipCode: "94102",
	}

	employee := Employee{
		Person:   Person{Name: "Frank", Age: 32, Email: "frank@company.com"},
		Address:  address,
		JobTitle: "Software Engineer",
		Salary:   120000,
	}

	fmt.Printf("   %+v\n", employee)
	fmt.Printf("   Name: %s\n", employee.Person.Name)
	fmt.Printf("   City: %s\n\n", employee.Address.City)

	// ===================================
	// 9. Struct Comparison
	// ===================================
	fmt.Println("9. Struct comparison:")
	person1 := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
	person2 := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
	person3 := Person{Name: "Bob", Age: 30, Email: "bob@example.com"}

	fmt.Printf("   person1 == person2: %t\n", person1 == person2)
	fmt.Printf("   person1 == person3: %t\n\n", person1 == person3)

	// ===================================
	// 10. Anonymous Structs
	// ===================================
	fmt.Println("10. Anonymous structs:")

	// Useful for temporary data structures
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}

	fmt.Printf("    Server: %s:%d\n\n", config.Host, config.Port)

	// ===================================
	// 11. Struct Tags (for JSON, DB, etc.)
	// ===================================
	fmt.Println("11. Struct tags:")

	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email,omitempty"`
	}

	user := User{ID: 1, Username: "alice", Email: "alice@example.com"}
	fmt.Printf("    User struct: %+v\n", user)
	fmt.Println("    (Tags used for JSON marshaling/unmarshaling)")

	// ===================================
	// 12. Struct Slices
	// ===================================
	fmt.Println("\n12. Slice of structs:")

	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@example.com"},
		{Name: "Bob", Age: 30, Email: "bob@example.com"},
		{Name: "Carol", Age: 28, Email: "carol@example.com"},
	}

	for i, person := range people {
		fmt.Printf("    %d: %s (%d)\n", i, person.Name, person.Age)
	}

	fmt.Println("\n=== Structs Tutorial Complete! ===")
}
