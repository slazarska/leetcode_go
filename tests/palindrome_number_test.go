package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected bool
	}{
		{"positive test", 101, true},
		{"negative test", 100, false},
		{"negative number", -101, false},
		{"zero", 0, true},
		{"minimum", -231, false},
		{"maximum", 230, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.IsPalindrome(tt.x)
			if result != tt.expected {
				t.Errorf("%s: IsPalindrome(%d) is %v, expected %v", tt.name, tt.x, result, tt.expected)
			}
		})
	}
}
