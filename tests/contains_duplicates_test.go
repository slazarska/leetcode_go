package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected bool
	}{
		{"positive case", []int{1, 2, 3, 1}, 3, true},
		{"negative case", []int{1, 2, 3}, 4, false},
		{"one int in array", []int{1}, 1, false},
		{"k is zero", []int{1, 2, 3, 1}, 0, false},
		{"all elements are the same", []int{2, 2, 2, 2, 2}, 2, true},
		{"large k", []int{1, 2, 3, 1, 2, 3}, 100000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.ContainsNearbyDuplicate(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("%s: ContainsNearbyDuplicate(%v, %d) = %v, expected %v", tt.name, tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
