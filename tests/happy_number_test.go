package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{19, true},          // счастливое число
		{2, false},          // несчастливое число
		{1, true},           // граничный случай
		{4, false},          // число с циклом
		{2147483647, false}, // большое число
		{16, false},         // число, которое превращается в само себя
		{33, false},         // Число с одинаковыми цифрами
	}

	for _, tt := range tests {
		result := easy.IsHappy(tt.num)
		if result != tt.expected {
			t.Errorf("IsHappy(%d) = %v; expected %v", tt.num, result, tt.expected)
		}
	}
}

func TestHappyNumberFloydsCycle(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{19, true},          // счастливое число
		{2, false},          // несчастливое число
		{1, true},           // граничный случай
		{4, false},          // число с циклом
		{2147483647, false}, // большое число
		{16, false},         // число, которое превращается в само себя
		{33, false},         // Число с одинаковыми цифрами
	}

	for _, tt := range tests {
		result := easy.IsHappyFloydsCycle(tt.num)
		if result != tt.expected {
			t.Errorf("IsHappyFloydsCycle(%d) = %v; expected %v", tt.num, result, tt.expected)
		}
	}
}
