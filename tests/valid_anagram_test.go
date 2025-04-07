package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s, t     string
		expected bool
	}{
		{"positive test", "tom marvolo riddle", "i am lord voldemort", true},
		{"negative test", "abc", "aaa", false},
		{"one letter", "a", "a", true},
		{"different lengths", "abc", "ab", false},
		{"unicode letters", "кот", "ток", true},
		{"emoji", "😺🐶", "🐶😺", true},
		{"empty strings", "", "", true},
		{"case insensitive", "Listen", "Silent", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.IsAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("%s: IsAnagram(%s, %s) = %v, expected %v", tt.name, tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
