package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Value receiver - cannot modify original
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Pointer receiver - can modify original
func (p *Person) HaveBirthday() {
	p.Age++
}

// Pointer receiver
func (p *Person) Rename(newName string) {
	p.Name = newName
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	fmt.Println("=== Pointers Challenge Solution ===\n")

	// ===================================
	// 1. Swap Function
	// ===================================
	fmt.Println("1. Swap function:")
	x, y := 10, 20
	fmt.Printf("   Before swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("   After swap: x=%d, y=%d\n\n", x, y)

	// ===================================
	// 2. Modify Person
	// ===================================
	fmt.Println("2. Modify person:")
	alice := Person{Name: "Alice", Age: 25}
	fmt.Printf("   Before: %s\n", alice.GetInfo())

	modifyPerson(&alice)
	fmt.Printf("   After modification: %s\n\n", alice.GetInfo())

	// ===================================
	// 3. Double Slice Values
	// ===================================
	fmt.Println("3. Double slice values:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Before: %v\n", numbers)

	doubleSliceValues(&numbers)
	fmt.Printf("   After doubling: %v\n\n", numbers)

	// ===================================
	// 4. Value vs Pointer Receivers
	// ===================================
	fmt.Println("4. Value vs pointer receivers:")
	bob := Person{Name: "Bob", Age: 30}

	fmt.Printf("   Initial: %s\n", bob.GetInfo())

	// Value receiver - doesn't modify original
	info := bob.GetInfo() // This uses value receiver
	fmt.Printf("   GetInfo(): %s\n", info)

	// Pointer receiver - modifies original
	bob.HaveBirthday()
	fmt.Printf("   After HaveBirthday(): %s\n", bob.GetInfo())

	bob.Rename("Robert")
	fmt.Printf("   After Rename(): %s\n\n", bob.GetInfo())

	// ===================================
	// 5. Persistence of Modifications
	// ===================================
	fmt.Println("5. Demonstrating persistence:")

	// Case A: Pass by value - doesn't persist
	carol := Person{Name: "Carol", Age: 28}
	fmt.Printf("   Before tryModifyByValue: %s\n", carol.GetInfo())
	tryModifyByValue(carol)
	fmt.Printf("   After tryModifyByValue: %s (unchanged)\n", carol.GetInfo())

	// Case B: Pass by pointer - persists
	fmt.Printf("   Before modifyByPointer: %s\n", carol.GetInfo())
	modifyByPointer(&carol)
	fmt.Printf("   After modifyByPointer: %s (changed)\n\n", carol.GetInfo())

	// ===================================
	// Bonus 1: Linked List
	// ===================================
	fmt.Println("Bonus 1: Pointer-based linked list:")
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}

	printList(head)
	fmt.Println()

	// ===================================
	// Bonus 2: Returning Pointers
	// ===================================
	fmt.Println("Bonus 2: Function returning pointer:")
	newPerson := createPerson("David", 35)
	fmt.Printf("   Created: %s\n", newPerson.GetInfo())

	// Can modify since we have pointer
	newPerson.HaveBirthday()
	fmt.Printf("   After birthday: %s\n\n", newPerson.GetInfo())

	// ===================================
	// Bonus 3: Nil Pointer Checking
	// ===================================
	fmt.Println("Bonus 3: Nil pointer checking:")
	var nilPerson *Person

	// Safe function that checks for nil
	safeModifyPerson(nilPerson)

	validPerson := &Person{Name: "Eve", Age: 27}
	safeModifyPerson(validPerson)
	fmt.Printf("   Valid person modified: %s\n", validPerson.GetInfo())

	// ===================================
	// Summary
	// ===================================
	fmt.Println("\n=== What You Learned ===")
	fmt.Println("✓ Using & (address-of) and * (dereference) operators")
	fmt.Println("✓ Passing by value vs passing by pointer")
	fmt.Println("✓ Value receivers vs pointer receivers")
	fmt.Println("✓ When modifications persist and when they don't")
	fmt.Println("✓ Pointer-based data structures")
	fmt.Println("✓ Returning pointers from functions")
	fmt.Println("✓ Nil pointer checking for safety")
}

// ===================================
// Required Functions
// ===================================

func swap(a, b *int) {
	*a, *b = *b, *a
}

func modifyPerson(p *Person) {
	p.Name = "Modified " + p.Name
	p.Age += 10
}

func doubleSliceValues(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

// ===================================
// Helper Functions
// ===================================

func tryModifyByValue(p Person) {
	p.Name = "Modified"
	p.Age = 100
	// Changes are lost when function returns
}

func modifyByPointer(p *Person) {
	p.Name = "Modified " + p.Name
	p.Age = 100
	// Changes persist
}

func printList(head *Node) {
	fmt.Print("   List: ")
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		if current.Next != nil {
			fmt.Print("-> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func createPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func safeModifyPerson(p *Person) {
	if p == nil {
		fmt.Println("   Cannot modify nil pointer")
		return
	}
	p.Age += 5
}
