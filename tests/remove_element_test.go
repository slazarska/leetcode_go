package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{"empty array", []int{}, 3, 0},
		{"all elements match", []int{2, 2, 2, 2}, 2, 0},
		{"no elements match", []int{1, 2, 3, 4, 5}, 6, 5},
		{"middle removal", []int{1, 2, 3, 3, 4, 5}, 3, 4},
		{"first and last removal", []int{3, 1, 2, 4, 3}, 3, 3},
	}

	for _, tt := range tests {
		result := easy.RemoveElement(tt.nums, tt.val)
		if result != tt.expected {
			t.Errorf("%s: RemoveElement(%v, %d) = %d; expected %d", tt.name, tt.nums, tt.val, result, tt.expected)
		}
	}
}
