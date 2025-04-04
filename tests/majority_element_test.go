package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"one element", []int{1}, 1},
		{"simple majority", []int{1, 2, 1}, 1},
		{"majority at start", []int{3, 3, 3, 1, 2}, 3},
		{"majority at end", []int{1, 2, 4, 5, 5, 5, 5}, 5},
		{"all same", []int{7, 7, 7, 7, 7}, 7},
		{"even length majority", []int{2, 2, 1, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.MajorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("%s: Majority Element(%v) = %d, expected: %d", tt.name, tt.nums, result, tt.expected)
			}
		})
	}
}
