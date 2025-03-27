package main

import (
	"fmt"
	"leetcode_go/easy"
)

func main() {
	//#2469. Конвертация температуры
	fmt.Println(easy.ConvertTemperature(22.00))

	//#88. Слияние отсортированных списков
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	easy.MergeSortedArrays(nums1, m, nums2, n)
	fmt.Println(nums1)
}
