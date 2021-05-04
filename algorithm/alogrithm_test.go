package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var array []int

func TestMain(m *testing.M) {
	fmt.Println("Before Testing")
	m.Run()
	fmt.Println("After Testing")
}

func TestTableAlgorithm(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{
			name:     "BubbleSort",
			array:    []int{6, 4, 5, 3, 7, 8, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BubbleSort(test.array)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkSub(b *testing.B) {
	array = []int{6, 5, 3, 1, 8, 7, 2, 4}
	b.Run("BubbleSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BubbleSort(array)
		}
	})
	b.Run("InsertionSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InsertionSort(array)
		}
	})
	b.Run("SelectionSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelectionSort(array)
		}
	})
	b.Run("ShellSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ShellSort(array)
		}
	})
}

func TestSub(t *testing.T) {
	array = []int{6, 5, 3, 1, 8, 7, 2, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Run("BubbleSort", func(t *testing.T) {
		result := BubbleSort(array)
		assert.Equal(t, expected, result)
	})
	t.Run("InsertionSort", func(t *testing.T) {
		result := InsertionSort(array)
		assert.Equal(t, expected, result)
	})
	t.Run("SelectionSort", func(t *testing.T) {
		result := SelectionSort(array)
		assert.Equal(t, expected, result)
	})
	t.Run("ShellSort", func(t *testing.T) {
		result := ShellSort(array)
		assert.Equal(t, expected, result)
	})
}
