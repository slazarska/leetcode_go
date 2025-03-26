package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected []float64
	}{
		{"Zero Celsius", 0, []float64{273.15, 32}},
		{"Boiling Point", 100, []float64{373.15, 212}},
		{"Freezing Point", -40, []float64{233.15, -40}},
		{"Room Temperature", 22, []float64{295.15, 71.6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := easy.ConvertTemperature(tt.celsius)
			if !AlmostEqual(result[0], tt.expected[0]) {
				t.Errorf("For Celsius %.2f, expected Kelvin %.2f, but got %.2f", tt.celsius, tt.expected[0], result[0])
			}

			if !AlmostEqual(result[1], tt.expected[1]) {
				t.Errorf("For Celsius %.2f, expected Fahrenheit %.2f, but got %.2f", tt.celsius, tt.expected[1], result[1])
			}
		})
	}
}
