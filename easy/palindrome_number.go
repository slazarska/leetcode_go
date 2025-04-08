package easy

/*
eng
#9. Palindrome Number
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:
-231 <= x <= 231 - 1

Follow up: Could you solve it without converting the integer to a string?
*/

/*
rus
#9. Число-палиндром
Для целого числа x вернуть true, если x является палиндромом, и false в противном случае.

Пример 1:
Входные данные: x = 121
Выходные данные: true
Пояснение: 121 читается как 121 слева направо и справа налево.

Пример 2:
Входные данные: x = -121
Выходные данные: false
Пояснение: Слева направо читается как -121. Справа налево становится 121-. Поэтому это не палиндром.

Пример 3:
Входные данные: x = 10
Выходные данные: false
Пояснение: Читается как 01 справа налево. Поэтому это не палиндром.

Ограничения:
-231 <= x <= 231 - 1

Продолжение: можете ли вы решить это, не преобразуя целое число в строку?
*/

func IsPalindrome(x int) bool {
	original := x
	reversed := 0

	for x > 0 {
		lastDigit := x % 10
		reversed = reversed*10 + lastDigit
		x /= 10
	}

	return original == reversed
}
