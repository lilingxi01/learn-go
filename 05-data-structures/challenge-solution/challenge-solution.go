package main

import (
	"fmt"
	"sort"
	"strings"
)

// Contact represents a contact in our address book
type Contact struct {
	Name  string
	Phone string
	Email string
}

func main() {
	fmt.Println("=== Contact Book Challenge Solution ===\n")

	// ===================================
	// Approach 1: Slice-Based Contact Book
	// ===================================
	fmt.Println("Approach 1: Slice-Based Implementation")
	fmt.Println("---------------------------------------")

	var contacts []Contact

	// Add contacts
	contacts = AddContact(contacts, Contact{"Alice Johnson", "555-0101", "alice@example.com"})
	contacts = AddContact(contacts, Contact{"Bob Smith", "555-0102", "bob@example.com"})
	contacts = AddContact(contacts, Contact{"Carol Williams", "555-0103", "carol@example.com"})
	contacts = AddContact(contacts, Contact{"David Brown", "555-0104", "david@example.com"})

	// List all contacts
	fmt.Println("\nAll Contacts:")
	ListContacts(contacts)

	// Find a contact
	fmt.Println("\nFinding 'Bob Smith':")
	if contact, found := FindContact(contacts, "Bob Smith"); found {
		fmt.Printf("Found: %s - %s - %s\n", contact.Name, contact.Phone, contact.Email)
	} else {
		fmt.Println("Contact not found")
	}

	// Delete a contact
	fmt.Println("\nDeleting 'Carol Williams':")
	contacts = DeleteContact(contacts, "Carol Williams")

	// List contacts after deletion
	fmt.Println("\nContacts After Deletion:")
	ListContacts(contacts)

	// Update a contact
	fmt.Println("\nUpdating Bob's email:")
	contacts = UpdateContact(contacts, "Bob Smith", Contact{
		Name:  "Bob Smith",
		Phone: "555-0102",
		Email: "bob.new@example.com",
	})
	ListContacts(contacts)

	// ===================================
	// Approach 2: Map-Based Contact Book (Bonus)
	// ===================================
	fmt.Println("\n\nApproach 2: Map-Based Implementation (Faster Lookups)")
	fmt.Println("------------------------------------------------------")

	contactBook := NewContactBook()

	// Add contacts
	contactBook.Add(Contact{"Alice Johnson", "555-0101", "alice@example.com"})
	contactBook.Add(Contact{"Bob Smith", "555-0102", "bob@example.com"})
	contactBook.Add(Contact{"Carol Williams", "555-0103", "carol@example.com"})
	contactBook.Add(Contact{"David Brown", "555-0104", "david@example.com"})

	// List all
	fmt.Println("\nAll Contacts (Map-Based):")
	contactBook.ListAll()

	// Find by name
	fmt.Println("\nFinding 'Alice Johnson':")
	if contact, found := contactBook.FindByName("Alice Johnson"); found {
		fmt.Printf("Found: %s - %s - %s\n", contact.Name, contact.Phone, contact.Email)
	}

	// Find by phone
	fmt.Println("\nFinding by phone '555-0102':")
	if contact, found := contactBook.FindByPhone("555-0102"); found {
		fmt.Printf("Found: %s - %s\n", contact.Name, contact.Phone)
	}

	// Delete
	fmt.Println("\nDeleting 'Bob Smith':")
	contactBook.Delete("Bob Smith")
	contactBook.ListAll()

	// ===================================
	// Approach 3: Sorted Contact Book
	// ===================================
	fmt.Println("\n\nApproach 3: Sorted Contact Book")
	fmt.Println("--------------------------------")

	sorted := []Contact{
		{"Zebra Smith", "555-9999", "zebra@example.com"},
		{"Alice Johnson", "555-0001", "alice@example.com"},
		{"Bob Williams", "555-0002", "bob@example.com"},
		{"Charlie Brown", "555-0003", "charlie@example.com"},
	}

	fmt.Println("\nBefore sorting:")
	for _, c := range sorted {
		fmt.Printf("  %s\n", c.Name)
	}

	// Sort alphabetically
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Name < sorted[j].Name
	})

	fmt.Println("\nAfter sorting:")
	for _, c := range sorted {
		fmt.Printf("  %s\n", c.Name)
	}

	fmt.Println("\n=== What You Learned ===")
	fmt.Println("✓ Creating and using structs")
	fmt.Println("✓ Slice operations (add, find, delete)")
	fmt.Println("✓ Map operations for fast lookups")
	fmt.Println("✓ Iterating over slices and maps")
	fmt.Println("✓ Returning multiple values (value, bool)")
	fmt.Println("✓ Sorting slices")
	fmt.Println("✓ Choosing right data structure for the task")
}

// ===================================
// Slice-Based Functions
// ===================================

func AddContact(contacts []Contact, contact Contact) []Contact {
	return append(contacts, contact)
}

func FindContact(contacts []Contact, name string) (Contact, bool) {
	for _, contact := range contacts {
		if contact.Name == name {
			return contact, true
		}
	}
	return Contact{}, false
}

func ListContacts(contacts []Contact) {
	if len(contacts) == 0 {
		fmt.Println("  No contacts")
		return
	}
	for i, contact := range contacts {
		fmt.Printf("  %d. %s - %s - %s\n", i+1, contact.Name, contact.Phone, contact.Email)
	}
}

func DeleteContact(contacts []Contact, name string) []Contact {
	for i, contact := range contacts {
		if contact.Name == name {
			// Remove by slicing
			return append(contacts[:i], contacts[i+1:]...)
		}
	}
	return contacts
}

func UpdateContact(contacts []Contact, name string, updated Contact) []Contact {
	for i, contact := range contacts {
		if contact.Name == name {
			contacts[i] = updated
			return contacts
		}
	}
	return contacts
}

// ===================================
// Map-Based Implementation (Bonus)
// ===================================

type ContactBook struct {
	contactsByName  map[string]Contact
	contactsByPhone map[string]Contact
}

func NewContactBook() *ContactBook {
	return &ContactBook{
		contactsByName:  make(map[string]Contact),
		contactsByPhone: make(map[string]Contact),
	}
}

func (cb *ContactBook) Add(contact Contact) {
	cb.contactsByName[contact.Name] = contact
	cb.contactsByPhone[contact.Phone] = contact
}

func (cb *ContactBook) FindByName(name string) (Contact, bool) {
	contact, found := cb.contactsByName[name]
	return contact, found
}

func (cb *ContactBook) FindByPhone(phone string) (Contact, bool) {
	contact, found := cb.contactsByPhone[phone]
	return contact, found
}

func (cb *ContactBook) Delete(name string) {
	if contact, found := cb.contactsByName[name]; found {
		delete(cb.contactsByName, name)
		delete(cb.contactsByPhone, contact.Phone)
	}
}

func (cb *ContactBook) ListAll() {
	if len(cb.contactsByName) == 0 {
		fmt.Println("  No contacts")
		return
	}

	// Get sorted names for consistent output
	names := make([]string, 0, len(cb.contactsByName))
	for name := range cb.contactsByName {
		names = append(names, name)
	}
	sort.Strings(names)

	for i, name := range names {
		contact := cb.contactsByName[name]
		fmt.Printf("  %d. %s - %s - %s\n", i+1, contact.Name, contact.Phone, contact.Email)
	}
}

func (cb *ContactBook) Search(query string) []Contact {
	results := []Contact{}
	query = strings.ToLower(query)

	for _, contact := range cb.contactsByName {
		if strings.Contains(strings.ToLower(contact.Name), query) ||
			strings.Contains(strings.ToLower(contact.Email), query) ||
			strings.Contains(contact.Phone, query) {
			results = append(results, contact)
		}
	}

	return results
}
