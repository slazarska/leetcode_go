package main

import (
	"fmt"
	"leetcode_go/easy"
)

func main() {

	//#9. Число-палиндром / Palindrome Number
	fmt.Println("#9. Palindrome Number:", easy.IsPalindrome(1001))

	//#26. Удалить дубликаты из отсортированного массива / Remove Duplicates From Sorted Array
	fmt.Println("#26. Remove Duplicates From Sorted Array:", easy.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 3}))

	//# 27. Удалить элемент / Remove element
	fmt.Println("#27. Remove Element:", easy.RemoveElement([]int{1, 2, 2, 3, 2}, 2))

	//#28. Индекс первого вхождения в строку / Find the Index of the First Occurrence in a String
	fmt.Println("#28. Find the Index of the First Occurrence in a String:", easy.FindIndexOfTheFirstOccurrenceInString("hello world", "world"))

	//#67. Бинарное сложение / Add Binary
	fmt.Println("#67. Add Binary:", easy.AddBinary("1010", "1011"))

	//#70. Подъем по лестнице / Climbing Stairs
	fmt.Println("#70. Climbing Stairs:", easy.ClimbStairs(5))

	//#88. Слияние отсортированных списков / Merge Sorted Arrays
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	easy.MergeSortedArrays(nums1, m, nums2, n)
	fmt.Println("#88. Merge Sorted Arrays:", nums1)

	//#121. Лучшее время для покупки и продажи акций / Best Time To Buy And Sell Stock
	fmt.Println("#121. Best Time To Buy And Sell Stock:", easy.MaxProfit([]int{7, 6, 3, 4, 5}))

	//#169. Элемент большинства / Majority Element
	fmt.Println("#169. Majority Element:", easy.MajorityElement([]int{0, 0, 1, 1, 2, 2}))

	//#202. Счастливое число / Happy Number
	fmt.Println("#202. Happy Number:", easy.IsHappy(19))

	//219. Содержит дубликаты II / Contains Dulicates II
	fmt.Println("#219. Contains Dulicates II:", easy.ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))

	//#242. Валидная анаграмма / Valid Anagram
	fmt.Println("#242. Valid Anagram:", easy.IsAnagram("tom marvolo riddle", "i am lord voldemort"))

	//#290. Словесный паттерн / Word Patter
	fmt.Println("#290. Word Pattern:", easy.WordPattern("ab", "dog cat"))

	//#2469. Конвертация температуры / Convert the Temperature
	fmt.Println("#2469. Convert the Temperature:", easy.ConvertTemperature(22.00))
}
