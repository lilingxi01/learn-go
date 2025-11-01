package calculator

import (
	"testing"
)

// ===================================
// Benchmark Tests
// ===================================

// BenchmarkAdd measures the performance of the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

// BenchmarkMultiply measures the performance of the Multiply function
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(100, 200)
	}
}

// BenchmarkSum measures the performance of Sum with different sizes
func BenchmarkSum(b *testing.B) {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer() // Reset timer after setup

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

// BenchmarkSumSmall benchmarks Sum with small slice
func BenchmarkSumSmall(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

// BenchmarkSumLarge benchmarks Sum with large slice
func BenchmarkSumLarge(b *testing.B) {
	numbers := make([]int, 10000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

// BenchmarkFactorial measures the performance of Factorial
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

// BenchmarkIsEven measures the performance of IsEven
func BenchmarkIsEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEven(12345)
	}
}

// BenchmarkAbs measures the performance of Abs
func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-12345)
	}
}

// ===================================
// Sub-benchmarks
// ===================================

func BenchmarkOperations(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(100, 200)
		}
	})

	b.Run("Multiply", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Multiply(100, 200)
		}
	})

	b.Run("Divide", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Divide(100, 2)
		}
	})
}
