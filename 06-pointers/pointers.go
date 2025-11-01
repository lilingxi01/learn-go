package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Value receiver - operates on a copy
func (p Person) GreetValue() string {
	return "Hello from " + p.Name
}

// Pointer receiver - can modify original
func (p *Person) HaveBirthday() {
	p.Age++
}

// Value receiver trying to modify (won't work)
func (p Person) TryToModify() {
	p.Age = 100 // Only modifies the copy
}

func main() {
	fmt.Println("=== Go Pointers Tutorial ===\n")

	// ===================================
	// 1. Basic Pointers
	// ===================================
	fmt.Println("1. Basic pointers:")
	x := 42
	p := &x // p is a pointer to x

	fmt.Printf("   x = %d\n", x)
	fmt.Printf("   p = %p (memory address)\n", p)
	fmt.Printf("   *p = %d (value at address)\n\n", *p)

	// ===================================
	// 2. Modifying via Pointer
	// ===================================
	fmt.Println("2. Modifying via pointer:")
	fmt.Printf("   Before: x = %d\n", x)
	*p = 100
	fmt.Printf("   After *p = 100: x = %d\n\n", x)

	// ===================================
	// 3. Passing by Value
	// ===================================
	fmt.Println("3. Passing by value (copies data):")
	num := 10
	fmt.Printf("   Before modifyValue: num = %d\n", num)
	modifyValue(num)
	fmt.Printf("   After modifyValue: num = %d (unchanged)\n\n", num)

	// ===================================
	// 4. Passing by Pointer
	// ===================================
	fmt.Println("4. Passing by pointer (modifies original):")
	fmt.Printf("   Before modifyPointer: num = %d\n", num)
	modifyPointer(&num)
	fmt.Printf("   After modifyPointer: num = %d (changed)\n\n", num)

	// ===================================
	// 5. Pointer to Struct
	// ===================================
	fmt.Println("5. Pointer to struct:")
	alice := Person{Name: "Alice", Age: 25}
	pAlice := &alice

	fmt.Printf("   Direct access: %s, %d\n", alice.Name, alice.Age)
	fmt.Printf("   Via pointer: %s, %d\n", pAlice.Name, pAlice.Age)
	fmt.Println("   (Go auto-dereferences struct pointers)")

	// ===================================
	// 6. Value Receiver vs Pointer Receiver
	// ===================================
	fmt.Println("\n6. Value receiver vs pointer receiver:")
	bob := Person{Name: "Bob", Age: 30}

	fmt.Printf("   Before birthday: %s is %d\n", bob.Name, bob.Age)
	bob.HaveBirthday() // Pointer receiver - modifies original
	fmt.Printf("   After birthday: %s is %d\n", bob.Name, bob.Age)

	bob.TryToModify() // Value receiver - only modifies copy
	fmt.Printf("   After TryToModify: %s is %d (unchanged)\n\n", bob.Name, bob.Age)

	// ===================================
	// 7. Creating Pointers with new
	// ===================================
	fmt.Println("7. Creating pointers with new:")
	pNum := new(int) // Creates a pointer to zero value
	fmt.Printf("   *pNum = %d (zero value)\n", *pNum)
	*pNum = 42
	fmt.Printf("   After assignment: *pNum = %d\n\n", *pNum)

	// ===================================
	// 8. Nil Pointers
	// ===================================
	fmt.Println("8. Nil pointers:")
	var pPerson *Person
	fmt.Printf("   pPerson == nil: %t\n", pPerson == nil)

	// Safe to check nil before using
	if pPerson != nil {
		fmt.Println("   Person name:", pPerson.Name)
	} else {
		fmt.Println("   Pointer is nil, can't access fields")
	}
	fmt.Println()

	// ===================================
	// 9. Pointer Arithmetic (Not Allowed!)
	// ===================================
	fmt.Println("9. Pointer arithmetic:")
	fmt.Println("   Go does NOT allow pointer arithmetic")
	fmt.Println("   This is a safety feature")
	fmt.Println("   Use slices for arrays/sequences\n")

	// ===================================
	// 10. Swap Function
	// ===================================
	fmt.Println("10. Swap function using pointers:")
	a, b := 10, 20
	fmt.Printf("    Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("    After swap: a=%d, b=%d\n\n", a, b)

	// ===================================
	// 11. Pointers in Slices and Maps
	// ===================================
	fmt.Println("11. Pointers in slices and maps:")
	people := []*Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Carol", Age: 28},
	}

	fmt.Println("    Before birthdays:")
	for _, person := range people {
		fmt.Printf("      %s: %d\n", person.Name, person.Age)
	}

	// Modify through pointers
	for _, person := range people {
		person.Age++
	}

	fmt.Println("    After birthdays:")
	for _, person := range people {
		fmt.Printf("      %s: %d\n", person.Name, person.Age)
	}

	// ===================================
	// 12. When to Use Pointers
	// ===================================
	fmt.Println("\n12. When to use pointers:")
	fmt.Println("    ✓ When you need to modify the value")
	fmt.Println("    ✓ For large structs (avoid copying)")
	fmt.Println("    ✓ When nil is a valid value")
	fmt.Println("    ✓ For sharing data between functions")
	fmt.Println("    ✗ Avoid for simple types (int, bool, string)")

	fmt.Println("\n=== Pointers Tutorial Complete! ===")
}

// Passing by value - receives a copy
func modifyValue(x int) {
	x = 999 // Only modifies the copy
}

// Passing by pointer - can modify original
func modifyPointer(x *int) {
	*x = 999 // Modifies the original
}

// Swap two integers using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}
