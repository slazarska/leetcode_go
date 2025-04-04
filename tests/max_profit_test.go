package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"got profit", []int{7, 6, 3, 4, 6}, 3},
		{"no profit", []int{7, 7, 6}, 0},
		{"single day", []int{5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.MaxProfit(tt.prices)
			if result != tt.expected {
				t.Errorf("%s: MaxProfit(%v) = %d, expected = %d", tt.name, tt.prices, result, tt.expected)
			}
		})
	}
}
