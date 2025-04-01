package easy

/*
eng
#67.Add Binary

Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"

Constraints:
1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
*/

/*
rus
#67.Бинарное сложение

Даны две бинарные строки a и b, вернуть их сумму как двоичную строку.

Пример 1:
Вход: a = "11", b = "1"
Выход: "100"

Пример 2:
Вход: a = "1010", b = "1011"
Выход: "10101"

Ограничения:
1 <= длина a, длина b <= 104
a и b состоят только из символов '0' или '1'.
Каждая строка не содержит начальных нулей, за исключением самого нуля.
*/

func AddBinary(a string, b string) string {
	//carry (перенос) — это переменная, которая будет хранить значение, передаваемое в следующий разряд. Изначально 0.
	carry := 0
	//result — строка, в которую будем накапливать результат.
	result := ""

	//i и j — это индексы последних символов строк a и b, так как сложение ведется справа налево
	i, j := len(a)-1, len(b)-1

	//Цикл продолжается, пока есть хотя бы один непосчитанный бит (i >= 0 || j >= 0) или carry > 0.
	for i >= 0 || j >= 0 || carry > 0 {
		//В sum сначала записываем carry, так как он влияет на сумму разряда.
		sum := carry
		//Если i не вышел за границы a, то прибавляем a[i] - '0' (это преобразование символа в число).
		if i >= 0 {
			sum += int(a[i] - '0')
			//Уменьшаем i, двигаясь влево.
			i--
		}
		//Аналогично, если j не вышел за границы b, прибавляем b[j] - '0' и уменьшаем j.
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		//sum % 2 дает текущий бит (0 или 1).
		//string(sum%2 + '0') преобразует число 0 или 1 в символ '0' или '1'.
		//Конкатенация + result добавляет этот бит в начало строки result.
		result = string(sum%2+'0') + result
		//Обновление переноса. Если sum больше 1, то carry становится 1, иначе 0.
		carry = sum / 2
	}
	return result
}
