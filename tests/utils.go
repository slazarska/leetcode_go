package tests

import "math"

// AlmostEqual функция для сравнения с учетом погрешности
func AlmostEqual(a float64, b float64) bool {
	const epsilon = 1e-5 //это погрешность = 0.000001
	return math.Abs(a-b) < epsilon
}
