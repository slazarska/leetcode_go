package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"minimum steps", 1, 1},
		{"2 steps", 2, 2},
		{"many steps", 15, 987},
		{"maximum steps", 45, 1836311903},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.ClimbStairs(tt.n)
			if result != tt.expected {
				t.Errorf("%s: ClimbStairs(%d) = %d, expected %d", tt.name, tt.n, result, tt.expected)
			}
		})
	}
}
