package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go Performance Optimization Examples ===\n")

	// ===================================
	// 1. String Building
	// ===================================
	fmt.Println("1. String concatenation performance:")

	// Inefficient
	start := time.Now()
	inefficientConcat()
	fmt.Printf("   Inefficient: %v\n", time.Since(start))

	// Efficient
	start = time.Now()
	efficientConcat()
	fmt.Printf("   Efficient (strings.Builder): %v\n\n", time.Since(start))

	// ===================================
	// 2. Slice Preallocation
	// ===================================
	fmt.Println("2. Slice allocation:")

	start = time.Now()
	withoutPrealloc()
	fmt.Printf("   Without prealloc: %v\n", time.Since(start))

	start = time.Now()
	withPrealloc()
	fmt.Printf("   With prealloc: %v\n\n", time.Since(start))

	fmt.Println("=== Performance Tutorial Complete! ===")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("✓ Use strings.Builder for string concatenation")
	fmt.Println("✓ Pre-allocate slices when size is known")
	fmt.Println("✓ Use sync.Pool for frequent allocations")
	fmt.Println("✓ Profile before optimizing")
	fmt.Println("✓ Benchmark to verify improvements")
}

func inefficientConcat() string {
	result := ""
	for i := 0; i < 10000; i++ {
		result += "a"
	}
	return result
}

func efficientConcat() string {
	var sb strings.Builder
	for i := 0; i < 10000; i++ {
		sb.WriteString("a")
	}
	return sb.String()
}

func withoutPrealloc() []int {
	var slice []int
	for i := 0; i < 100000; i++ {
		slice = append(slice, i)
	}
	return slice
}

func withPrealloc() []int {
	slice := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		slice = append(slice, i)
	}
	return slice
}
