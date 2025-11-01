package calculator

import (
	"testing"
)

// ===================================
// 1. Basic Tests
// ===================================

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}

// ===================================
// 2. Table-Driven Tests
// ===================================

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zeros", 0, 0, 0},
		{"with zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 3, 4, 12},
		{"negative", -3, 4, -12},
		{"both negative", -3, -4, 12},
		{"with zero", 5, 0, 0},
		{"with one", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// ===================================
// 3. Error Testing
// ===================================

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative dividend", -10, 2, -5, false},
		{"negative divisor", 10, -2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
				return
			}

			if err != nil {
				t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				return
			}

			if got != tt.want {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivideByZeroError(t *testing.T) {
	_, err := Divide(10, 0)
	if err != ErrDivisionByZero {
		t.Errorf("expected ErrDivisionByZero, got %v", err)
	}
}

// ===================================
// 4. Boolean Tests
// ===================================

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"even positive", 4, true},
		{"odd positive", 5, false},
		{"even negative", -4, true},
		{"odd negative", -5, false},
		{"zero", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.n)
			if got != tt.want {
				t.Errorf("IsEven(%d) = %t; want %t", tt.n, got, tt.want)
			}
		})
	}
}

// ===================================
// 5. Edge Case Tests
// ===================================

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
		{"large positive", 1000000, 1000000},
		{"large negative", -1000000, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs(tt.n)
			if got != tt.want {
				t.Errorf("Abs(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}

// ===================================
// 6. Slice Tests
// ===================================

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{"normal", []int{1, 2, 3, 4, 5}, 15},
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed signs", []int{-5, 10, -3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.numbers)
			if got != tt.want {
				t.Errorf("Sum(%v) = %d; want %d", tt.numbers, got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    float64
	}{
		{"normal", []int{1, 2, 3, 4, 5}, 3.0},
		{"empty slice", []int{}, 0.0},
		{"single element", []int{10}, 10.0},
		{"even count", []int{2, 4, 6, 8}, 5.0},
		{"odd count", []int{1, 2, 3}, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Average(tt.numbers)
			if got != tt.want {
				t.Errorf("Average(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
			}
		})
	}
}

// ===================================
// 7. Parallel Tests
// ===================================

func TestMaxParallel(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 5, 3, 5},
		{"both negative", -5, -3, -3},
		{"mixed", -5, 3, 3},
		{"equal", 5, 5, 5},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Run in parallel
			got := Max(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// ===================================
// 8. Helper Function
// ===================================

func assertEqual(t *testing.T, got, want int) {
	t.Helper() // Marks this as a helper function
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestMinWithHelper(t *testing.T) {
	assertEqual(t, Min(3, 5), 3)
	assertEqual(t, Min(-3, -5), -5)
	assertEqual(t, Min(0, 0), 0)
}

// ===================================
// 9. Test with Setup and Teardown
// ===================================

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{"zero", 0, 1, false},
		{"one", 1, 1, false},
		{"small", 5, 120, false},
		{"negative", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.n)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Factorial(%d) expected error, got nil", tt.n)
				}
				return
			}

			if err != nil {
				t.Errorf("Factorial(%d) unexpected error: %v", tt.n, err)
				return
			}

			if got != tt.want {
				t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
