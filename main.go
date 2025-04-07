package main

import (
	"fmt"
	"leetcode_go/easy"
)

func main() {

	//#26. Удалить дубликаты из отсортированного массива / Remove Duplicates From Sorted Array
	fmt.Println("#26. RemoveDuplicates:", easy.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 3}))

	//# 27. Удалить элемент / Remove element
	fmt.Println("#27. RemoveElement:", easy.RemoveElement([]int{1, 2, 2, 3, 2}, 2))

	//#28. Индекс первого вхождения в строку / Find the Index of the First Occurrence in a String
	fmt.Println("#28. FindIndexOfTheFirstOccurrenceInString:", easy.FindIndexOfTheFirstOccurrenceInString("hello world", "world"))

	//#67. Бинарное сложение / Add Binary
	fmt.Println("#67. AddBinary:", easy.AddBinary("1010", "1011"))

	//#88. Слияние отсортированных списков / Merge Sorted Arrays
	easy.MergeSortedArrays([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	//#121. Лучшее время для покупки и продажи акций / Best Time To Buy And Sell Stock
	fmt.Println("#121. MaxProfit:", easy.MaxProfit([]int{7, 6, 3, 4, 5}))

	//#169. Элемент большинства / Majority Element
	fmt.Println("#169. MajorityElement:", easy.MajorityElement([]int{0, 0, 1, 1, 2, 2}))

	//#202. Счастливое число / Happy Number
	fmt.Println("#202. IsHappy:", easy.IsHappy(19))

	//219. Содержит дубликаты II / Contains Dulicates II
	fmt.Println("#219. ContainsNearbyDuplicate:", easy.ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))

	//#242. Валидная анаграмма / Valid Anagram
	fmt.Println("#242. IsAnagram:", easy.IsAnagram("tom marvolo riddle", "i am lord voldemort"))

	//#290. Словесный паттерн / Word Patter
	fmt.Println("#290. WordPattern:", easy.WordPattern("ab", "dog cat"))

	//#2469. Конвертация температуры / Convert the Temperature
	fmt.Println("#2469. ConvertTemperature:", easy.ConvertTemperature(22.00))
}
