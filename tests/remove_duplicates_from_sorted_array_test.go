package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"multiple unique", []int{0, 1, 1, 2, 2, 3}, 4},
		{"already unique", []int{1, 2, 3, 4, 5}, 5},
		{"single element", []int{7}, 1},
	}

	for _, tt := range tests {
		result := easy.RemoveDuplicates(tt.nums)
		if result != tt.expected {
			t.Errorf("%s: RemoveDuplicatesFromSortedArray(%v) = %d, expected = %d", tt.name, tt.nums, result, tt.expected)
		}
	}
}
