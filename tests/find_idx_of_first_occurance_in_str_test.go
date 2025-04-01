package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestFindIndexOfTheFirstOccurrenceInString(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{"needle in the middle", "hello world", "world", 6},
		{"needle at the beginning", "hello world", "hello", 0},
		{"empty needle", "hello world", "", 0},
		{"empty haystack", "", "world", -1},
		{"both empty", "", "", 0},
		{"needle not found", "hello world", "x", -1},
		{"multiple occurrences", "hello hello world", "hello", 0},
		{"needle longer than haystrack", "abc", "abcd", -1},
	}

	for _, tt := range tests {
		t.Run(tt.haystack+"_"+tt.needle, func(t *testing.T) {
			got := easy.FindIndexOfTheFirstOccurrenceInString(tt.haystack, tt.needle)
			if got != tt.expected {
				t.Errorf("FindIndexOfTheFirstOccurrenceInString(%q, %q) = %d; expected %d", tt.haystack, tt.needle, got, tt.expected)
			}
		})
	}
}
