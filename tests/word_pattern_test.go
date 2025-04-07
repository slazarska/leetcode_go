package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name       string
		pattern, s string
		expected   bool
	}{
		{"positive test", "aba", "dog cat dog", true},
		{"negative test", "aba", "dog cat rabbit", false},
		{"pattern is shorter than string", "a", "dog cat", false},
		{"pattern is longer than string", "abab", "dog", false},
		{"same word maps to different letters", "ab", "dog dog", false},
		{"same letter maps to different words", "aa", "dog cat", false},
		{"all same letters and words", "aaa", "dog dog dog", true},
		{"one letter, one word", "a", "dog", true},
		{"different letters, same word", "abc", "dog dog dog", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.WordPattern(tt.pattern, tt.s)
			if result != tt.expected {
				t.Errorf("%s: WordPattern(%s, %s) = %v, expected %v", tt.name, tt.pattern, tt.s, result, tt.expected)
			}
		})
	}
}
