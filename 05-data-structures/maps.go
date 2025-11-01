package main

import "fmt"

func main() {
	fmt.Println("=== Go Maps Tutorial ===\n")

	// ===================================
	// 1. Creating Maps
	// ===================================
	fmt.Println("1. Creating maps:")

	// Using make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Printf("   Ages: %v\n", ages)

	// Map literal
	capitals := map[string]string{
		"USA":    "Washington, D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Printf("   Capitals: %v\n\n", capitals)

	// ===================================
	// 2. Accessing Map Values
	// ===================================
	fmt.Println("2. Accessing map values:")

	// Direct access
	fmt.Printf("   Alice's age: %d\n", ages["Alice"])

	// Check if key exists
	age, exists := ages["Carol"]
	if exists {
		fmt.Printf("   Carol's age: %d\n", age)
	} else {
		fmt.Println("   Carol not found in map")
	}
	fmt.Println()

	// ===================================
	// 3. Adding and Updating
	// ===================================
	fmt.Println("3. Adding and updating:")
	scores := make(map[string]int)

	scores["Alice"] = 95
	fmt.Printf("   Added Alice: %v\n", scores)

	scores["Bob"] = 87
	scores["Carol"] = 92
	fmt.Printf("   Added Bob and Carol: %v\n", scores)

	scores["Alice"] = 98 // Update
	fmt.Printf("   Updated Alice's score: %v\n\n", scores)

	// ===================================
	// 4. Deleting Entries
	// ===================================
	fmt.Println("4. Deleting entries:")
	fmt.Printf("   Before delete: %v\n", scores)

	delete(scores, "Bob")
	fmt.Printf("   After deleting Bob: %v\n\n", scores)

	// ===================================
	// 5. Iterating Over Maps
	// ===================================
	fmt.Println("5. Iterating over maps:")
	for name, score := range scores {
		fmt.Printf("   %s: %d\n", name, score)
	}
	fmt.Println()

	// ===================================
	// 6. Map Length
	// ===================================
	fmt.Println("6. Map length:")
	fmt.Printf("   Number of entries: %d\n\n", len(scores))

	// ===================================
	// 7. Map of Slices
	// ===================================
	fmt.Println("7. Map of slices:")
	grades := make(map[string][]int)

	grades["Alice"] = []int{95, 87, 92}
	grades["Bob"] = []int{78, 85, 90}

	for student, gradeList := range grades {
		fmt.Printf("   %s: %v\n", student, gradeList)
	}
	fmt.Println()

	// ===================================
	// 8. Map of Maps
	// ===================================
	fmt.Println("8. Nested maps:")
	users := make(map[string]map[string]string)

	users["alice"] = make(map[string]string)
	users["alice"]["name"] = "Alice Smith"
	users["alice"]["email"] = "alice@example.com"

	users["bob"] = map[string]string{
		"name":  "Bob Johnson",
		"email": "bob@example.com",
	}

	for username, info := range users {
		fmt.Printf("   %s: %v\n", username, info)
	}
	fmt.Println()

	// ===================================
	// 9. Map as Set
	// ===================================
	fmt.Println("9. Using map as a set:")
	set := make(map[string]bool)

	// Add items
	set["apple"] = true
	set["banana"] = true
	set["cherry"] = true

	// Check membership
	if set["apple"] {
		fmt.Println("   Set contains 'apple'")
	}
	if !set["grape"] {
		fmt.Println("   Set does not contain 'grape'")
	}

	// List all items
	fmt.Print("   Set items: ")
	for item := range set {
		fmt.Printf("%s ", item)
	}
	fmt.Println("\n")

	// ===================================
	// 10. Counting with Maps
	// ===================================
	fmt.Println("10. Counting with maps:")
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++ // Increments even if key doesn't exist
	}

	for word, count := range wordCount {
		fmt.Printf("    %s: %d\n", word, count)
	}
	fmt.Println()

	// ===================================
	// 11. Grouping with Maps
	// ===================================
	fmt.Println("11. Grouping data with maps:")

	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Carol", 25},
		{"Dave", 30},
		{"Eve", 25},
	}

	byAge := make(map[int][]string)
	for _, p := range people {
		byAge[p.Age] = append(byAge[p.Age], p.Name)
	}

	for age, names := range byAge {
		fmt.Printf("    Age %d: %v\n", age, names)
	}
	fmt.Println()

	// ===================================
	// 12. Zero Values and Safe Access
	// ===================================
	fmt.Println("12. Zero values and safe access:")
	var nilMap map[string]int // nil map

	// Reading from nil map is safe (returns zero value)
	fmt.Printf("    Reading from nil map: %d\n", nilMap["key"])

	// But writing to nil map causes panic!
	// nilMap["key"] = 1 // This would panic!

	// Always initialize with make
	safeMap := make(map[string]int)
	safeMap["key"] = 1
	fmt.Printf("    Safe map after initialization: %v\n", safeMap)

	fmt.Println("\n=== Maps Tutorial Complete! ===")
}
