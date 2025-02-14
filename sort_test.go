package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "Duplicate elements",
			input:    []int{3, 3, 3, 1, 1, 2},
			expected: []int{1, 1, 2, 3, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "Duplicate elements",
			input:    []int{3, 3, 3, 1, 1, 2},
			expected: []int{1, 1, 2, 3, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			InsertionSort(input)
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("InsertionSort() = %v, want %v", input, tt.expected)
			}
		})
	}
}

// Benchmark tests to compare performance
func BenchmarkMergeSort(b *testing.B) {
	input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	for i := 0; i < b.N; i++ {
		MergeSort(input)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	for i := 0; i < b.N; i++ {
		testInput := make([]int, len(input))
		copy(testInput, input)
		InsertionSort(testInput)
	}
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}

func BenchmarkSortSmall(b *testing.B) {
	input := generateRandomSlice(100)

	b.Run("MergeSort-100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			MergeSort(testSlice)
		}
	})

	b.Run("InsertionSort-100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			InsertionSort(testSlice)
		}
	})
}

func BenchmarkSortMedium(b *testing.B) {
	input := generateRandomSlice(1000)

	b.Run("MergeSort-1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			MergeSort(testSlice)
		}
	})

	b.Run("InsertionSort-1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			InsertionSort(testSlice)
		}
	})
}

func BenchmarkSortLarge(b *testing.B) {
	input := generateRandomSlice(10000)

	b.Run("MergeSort-10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			MergeSort(testSlice)
		}
	})

	b.Run("InsertionSort-10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			InsertionSort(testSlice)
		}
	})
}

func BenchmarkSortVeryLarge(b *testing.B) {
	input := generateRandomSlice(100000)

	b.Run("MergeSort-100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			MergeSort(testSlice)
		}
	})

	b.Run("InsertionSort-100000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testSlice := make([]int, len(input))
			copy(testSlice, input)
			InsertionSort(testSlice)
		}
	})
}
