package easy

import "strings"

/*
eng
#28. Find the Index of the First Occurrence in a String

Given two strings needle and haystack, return the index of the first occurrence of needle
in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

Constraints:
1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
*/

/*
rus
#28. Найдите индекс первого вхождения в строку

Данные две строки needle и haystack, верните индекс первого вхождения needle
в haystack или -1, если needle не является частью haystack.

Пример 1:
Входные данные: haystack = "sadbutsad", needle = "sad"
Выходные данные: 0
Пояснение: "sad" встречается в индексах 0 и 6.
Первое вхождение встречается в индексе 0, поэтому мы возвращаем 0.

Пример 2:
Входные данные: haystack = "leetcode", needle = "leeto"
Выходные данные: -1
Пояснение: "leeto" не встречается в "leetcode", поэтому мы возвращаем -1.

Ограничения:
1 <= длина haystack, длина needle <= 104
haystack и needle состоят только из строчных английских символов.
*/

func FindIndexOfTheFirstOccurrenceInString(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
