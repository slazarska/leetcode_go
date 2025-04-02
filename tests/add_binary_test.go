package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b     string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"111", "1", "1000"},
	}

	for _, tt := range tests {
		result := easy.AddBinary(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("AddBinary(%s, %s) = %s; expected %s", tt.a, tt.b, result, tt.expected)
		}
	}
}
