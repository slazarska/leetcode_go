package easy

import "strings"

/*
eng
#242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

/*
rus
#242. Валидная анаграмма

Даны две строки s и t, верните true, если t является анаграммой s, и false в противном случае.

Пример 1:
Входные данные: s = "anagram", t = "nagaram"
Выходные данные: true

Пример 2:
Входные данные: s = "rat", t = "car"
Выходные данные: false

Ограничения:
1 <= длина s, длина t <= 5 * 104
s и t состоят из строчных английских букв.

Дополнение: Что, если входные данные содержат символы Unicode? Как бы вы адаптировали свое решение к такому случаю?
*/

func IsAnagram(s string, t string) bool {
	s = strings.ReplaceAll(s, " ", "")
	t = strings.ReplaceAll(t, " ", "")

	//Если строки разной длины, они не могут быть анаграммами.
	if len(s) != len(t) {
		return false
	}

	//Создадим мапу, чтобы посчитать количество вхождений каждого символа.
	//Сразу возьмем руны и покроем кейс с Unicode.
	count := make(map[rune]int)

	//Считаем символы в первой строке, увеличия счетчик
	for _, char := range s {
		count[char]++
	}

	//Во второй строке наоборот счетчик символов уменьшаем
	for _, char := range t {
		count[char]--
		//Если счетчик стал отрицательным, значит, символа в t больше, чем в s, значит, это не анаграмма
		if count[char] < 0 {
			return false
		}
	}

	//Если все счётчики стали нулями, значит строки полностью одинаковые, значит это анаграммы.
	return true
}
