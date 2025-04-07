package easy

import "strings"

/*
eng
#290. Word Pattern

Given a pattern and a string s, find if s follows the same pattern.
Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
Specifically:
Each letter in pattern maps to exactly one unique word in s.
Each unique word in s maps to exactly one letter in pattern.
No two letters map to the same word, and no two words map to the same letter.


Example 1:
Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Explanation:
The bijection can be established as:
'a' maps to "dog".
'b' maps to "cat".

Example 2:
Input: pattern = "abba", s = "dog cat cat fish"
Output: false

Example 3:
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false

Constraints:
1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.
*/

/*
rus
#290. Словесный паттерн

Даны паттерн и строка s. Определите, следует ли s данному паттерну.
Здесь следует означает полное совпадение, такое, что существует биекция	 между буквой в паттерне и непустым словом в s.
А именно:
Каждая буква в паттерне маппится на ровно на одно уникальное слово в s.
Каждое уникальное слово в s маппится ровно на одну букву в паттерне.
Никакие две буквы не маппятся на одно и то же слово, и никакие два слова не маппятся на одну и ту же букву.

Пример 1:
Входные данные: pattern = "abba", s = "dog cat cat dog"
Выходные данные: true
Объяснение:
Биекция может быть установлена следующим образом:
'a' отображается на "dog".
'b' отображается на "cat".

Пример 2:
Вход: pattern = "abba", s = "dog cat cat fish"
Выход: false

Пример 3:
Вход: pattern = "aaaa", s = "dog cat cat dog"
Выход: false

Ограничения:
1 <= длина паттерна <= 300
паттерн содержит только строчные английские буквы.
1 <= длина строки s <= 3000
s содержит только строчные английские буквы и пробелы ' '.
s не содержит начальных или конечных пробелов.
Все слова в s разделены одним пробелом.
*/

func WordPattern(pattern string, s string) bool {
	//Вычленяем из строки отдельные слова
	words := strings.Split(s, " ")

	//Сравниваем число букв в паттерне с числом слов и если не совпадает, сразу возвращаем false
	if len(pattern) != len(words) {
		return false
	}

	//Создаем 2 мапы, первая хранит маппинг букв на слова, вторая слов на буквы
	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)

	//Создаем цикл по паттерну
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]

		//Если буква уже есть в мапе и связана с другим словом, то возвращаем false
		if mappedWord, ok := patternToWord[p]; ok {
			if mappedWord != w {
				return false
			}
			//Если не связана - то устанавливаем маппинг
		} else {
			patternToWord[p] = w
		}

		//Аналогично со второй мапой
		if mappedPat, ok := wordToPattern[w]; ok {
			if mappedPat != p {
				return false
			}
		} else {
			wordToPattern[w] = p
		}
	}

	return true
}
