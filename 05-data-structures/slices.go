package main

import "fmt"

func main() {
	fmt.Println("=== Go Slices Tutorial ===\n")

	// ===================================
	// 1. Creating Slices
	// ===================================
	fmt.Println("1. Creating slices:")

	// Slice literal
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Printf("   Fruits: %v\n", fruits)

	// Using make
	numbers := make([]int, 5) // length 5, capacity 5
	fmt.Printf("   Numbers (make): %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))

	// Make with capacity
	scores := make([]int, 3, 10) // length 3, capacity 10
	fmt.Printf("   Scores: %v, len=%d, cap=%d\n\n", scores, len(scores), cap(scores))

	// ===================================
	// 2. Append Operation
	// ===================================
	fmt.Println("2. Append operation:")

	var items []int
	fmt.Printf("   Initial: %v, len=%d, cap=%d\n", items, len(items), cap(items))

	items = append(items, 1)
	fmt.Printf("   After append(1): %v, len=%d, cap=%d\n", items, len(items), cap(items))

	items = append(items, 2, 3, 4)
	fmt.Printf("   After append(2,3,4): %v, len=%d, cap=%d\n\n", items, len(items), cap(items))

	// ===================================
	// 3. Slicing Operations
	// ===================================
	fmt.Println("3. Slicing operations:")
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("   Original: %v\n", data)
	fmt.Printf("   data[2:5]: %v\n", data[2:5]) // Elements 2, 3, 4
	fmt.Printf("   data[:4]: %v\n", data[:4])   // Elements 0 to 3
	fmt.Printf("   data[6:]: %v\n", data[6:])   // Elements 6 to end
	fmt.Printf("   data[:]: %v\n\n", data[:])   // All elements

	// ===================================
	// 4. Copy Operation
	// ===================================
	fmt.Println("4. Copy operation:")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("   Source: %v\n", src)
	fmt.Printf("   Destination: %v\n", dst)
	fmt.Printf("   Copied %d elements\n\n", copied)

	// ===================================
	// 5. Iterating Over Slices
	// ===================================
	fmt.Println("5. Iterating over slices:")
	colors := []string{"red", "green", "blue"}

	for i, color := range colors {
		fmt.Printf("   Index %d: %s\n", i, color)
	}
	fmt.Println()

	// ===================================
	// 6. Multidimensional Slices
	// ===================================
	fmt.Println("6. Multidimensional slices (2D):")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i, row := range matrix {
		fmt.Printf("   Row %d: %v\n", i, row)
	}
	fmt.Println()

	// ===================================
	// 7. Removing Elements
	// ===================================
	fmt.Println("7. Removing elements:")
	nums := []int{10, 20, 30, 40, 50}
	fmt.Printf("   Original: %v\n", nums)

	// Remove element at index 2
	removeIndex := 2
	nums = append(nums[:removeIndex], nums[removeIndex+1:]...)
	fmt.Printf("   After removing index %d: %v\n\n", removeIndex, nums)

	// ===================================
	// 8. Inserting Elements
	// ===================================
	fmt.Println("8. Inserting elements:")
	values := []int{1, 2, 4, 5}
	fmt.Printf("   Original: %v\n", values)

	// Insert 3 at index 2
	insertIndex := 2
	insertValue := 3
	values = append(values[:insertIndex], append([]int{insertValue}, values[insertIndex:]...)...)
	fmt.Printf("   After inserting %d at index %d: %v\n\n", insertValue, insertIndex, values)

	// ===================================
	// 9. Slice Capacity Growth
	// ===================================
	fmt.Println("9. Slice capacity growth:")
	var dynamic []int
	for i := 0; i < 10; i++ {
		dynamic = append(dynamic, i)
		fmt.Printf("   len=%d, cap=%d, data=%v\n", len(dynamic), cap(dynamic), dynamic)
	}
	fmt.Println()

	// ===================================
	// 10. Filtering Slices
	// ===================================
	fmt.Println("10. Filtering slices:")
	allNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNums := []int{}

	for _, num := range allNums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}

	fmt.Printf("    Original: %v\n", allNums)
	fmt.Printf("    Even numbers: %v\n\n", evenNums)

	// ===================================
	// 11. Slice References
	// ===================================
	fmt.Println("11. Slice references (slices share underlying array):")
	original := []int{1, 2, 3, 4, 5}
	reference := original[1:4]

	fmt.Printf("    Original: %v\n", original)
	fmt.Printf("    Reference: %v\n", reference)

	reference[0] = 99 // Modifies original!
	fmt.Printf("    After modifying reference[0]=99:\n")
	fmt.Printf("    Original: %v\n", original)
	fmt.Printf("    Reference: %v\n\n", reference)

	// ===================================
	// 12. Common Slice Patterns
	// ===================================
	fmt.Println("12. Common slice patterns:")

	// Reverse slice
	rev := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	fmt.Printf("    Reversed: %v\n", rev)

	// Check if slice contains element
	needle := 3
	haystack := []int{1, 2, 3, 4, 5}
	found := false
	for _, v := range haystack {
		if v == needle {
			found = true
			break
		}
	}
	fmt.Printf("    %v contains %d: %t\n", haystack, needle, found)

	fmt.Println("\n=== Slices Tutorial Complete! ===")
}
