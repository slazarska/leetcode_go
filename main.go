package main

import (
	"fmt"
	"leetcode_go/easy"
)

func main() {

	//# 27. Удалить элемент / Remove element
	fmt.Println(easy.RemoveElement([]int{1, 2, 2, 3, 2}, 2))

	//#28. Индекс первого вхождения в строку / Find the Index of the First Occurrence in a String
	fmt.Println(easy.FindIndexOfTheFirstOccurrenceInString("hello world", "world"))

	//#67. Бинарное сложение / Add Binary
	a := "1010"
	b := "1011"
	fmt.Println(easy.AddBinary(a, b))

	//#88. Слияние отсортированных списков / Merge Sorted Arrays
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	easy.MergeSortedArrays(nums1, m, nums2, n)
	fmt.Println(nums1)

	//#202. Счастливое число / Happy Number
	fmt.Println(easy.IsHappy(19))
	fmt.Println(easy.IsHappyFloydsCycle(16))

	//#2469. Конвертация температуры / Convert the Temperature
	fmt.Println(easy.ConvertTemperature(22.00))
}
