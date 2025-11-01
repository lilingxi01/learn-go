package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person represents a person with name and age
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // omitempty: exclude if empty
}

// Book represents a book with metadata
type Book struct {
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Year    int      `json:"year"`
	Pages   int      `json:"pages"`
	Tags    []string `json:"tags"`
	InPrint bool     `json:"in_print"`
}

func main() {
	fmt.Println("=== Go JSON Operations Tutorial ===\n")

	// ===================================
	// 1. Marshal (Encode) to JSON
	// ===================================
	fmt.Println("1. Marshaling struct to JSON:")

	person := Person{
		Name:  "Alice Johnson",
		Age:   30,
		Email: "alice@example.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   JSON: %s\n\n", string(jsonData))

	// ===================================
	// 2. Marshal with Indentation
	// ===================================
	fmt.Println("2. Pretty-print JSON (with indentation):")

	jsonPretty, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   JSON:\n%s\n\n", string(jsonPretty))

	// ===================================
	// 3. Unmarshal (Decode) from JSON
	// ===================================
	fmt.Println("3. Unmarshaling JSON to struct:")

	jsonStr := `{"name":"Bob Smith","age":25,"email":"bob@example.com"}`

	var newPerson Person
	err = json.Unmarshal([]byte(jsonStr), &newPerson)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   Decoded: %+v\n\n", newPerson)

	// ===================================
	// 4. Working with Slices
	// ===================================
	fmt.Println("4. JSON with slices:")

	people := []Person{
		{"Alice", 30, "alice@example.com"},
		{"Bob", 25, "bob@example.com"},
		{"Carol", 28, "carol@example.com"},
	}

	peopleJSON, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   JSON Array:\n%s\n\n", string(peopleJSON))

	// ===================================
	// 5. Nested Structures
	// ===================================
	fmt.Println("5. Nested JSON structures:")

	book := Book{
		Title:   "The Go Programming Language",
		Author:  "Donovan & Kernighan",
		Year:    2015,
		Pages:   400,
		Tags:    []string{"programming", "go", "tutorial"},
		InPrint: true,
	}

	bookJSON, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   Book JSON:\n%s\n\n", string(bookJSON))

	// ===================================
	// 6. Write JSON to File
	// ===================================
	fmt.Println("6. Writing JSON to file:")

	file, err := os.Create("data.json")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(people)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Println("   ✓ JSON written to data.json\n")

	// ===================================
	// 7. Read JSON from File
	// ===================================
	fmt.Println("7. Reading JSON from file:")

	readFile, err := os.Open("data.json")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer readFile.Close()

	var readPeople []Person
	decoder := json.NewDecoder(readFile)
	err = decoder.Decode(&readPeople)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   Read %d people from file:\n", len(readPeople))
	for _, p := range readPeople {
		fmt.Printf("   - %s (%d)\n", p.Name, p.Age)
	}
	fmt.Println()

	// ===================================
	// 8. Working with Maps
	// ===================================
	fmt.Println("8. JSON with maps:")

	dataMap := map[string]interface{}{
		"name":   "Dynamic Data",
		"count":  42,
		"active": true,
		"tags":   []string{"go", "json"},
	}

	mapJSON, err := json.MarshalIndent(dataMap, "", "  ")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   Map as JSON:\n%s\n\n", string(mapJSON))

	// ===================================
	// 9. Omitempty Tag
	// ===================================
	fmt.Println("9. omitempty tag behavior:")

	personWithEmail := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	personNoEmail := Person{Name: "Bob", Age: 25, Email: ""}

	// Safe to ignore error for known-good structs
	withEmail, _ := json.Marshal(personWithEmail)
	noEmail, _ := json.Marshal(personNoEmail)

	fmt.Printf("   With email: %s\n", string(withEmail))
	fmt.Printf("   No email: %s (email field omitted)\n\n", string(noEmail))

	// ===================================
	// 10. Handling JSON Errors
	// ===================================
	fmt.Println("10. Handling invalid JSON:")

	invalidJSON := `{"name": "Invalid", "age": "not a number"}`

	var invalidPerson Person
	err = json.Unmarshal([]byte(invalidJSON), &invalidPerson)
	if err != nil {
		fmt.Printf("    Error parsing invalid JSON: %v\n", err)
	}
	fmt.Println()

	// ===================================
	// Cleanup
	// ===================================
	os.Remove("data.json")

	fmt.Println("=== JSON Tutorial Complete! ===")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("✓ json.Marshal for encoding")
	fmt.Println("✓ json.Unmarshal for decoding")
	fmt.Println("✓ Use struct tags for field mapping")
	fmt.Println("✓ json.Encoder/Decoder for streams")
	fmt.Println("✓ MarshalIndent for pretty-printing")
	fmt.Println("✓ omitempty for optional fields")
}
