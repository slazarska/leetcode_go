package easy

/*
eng

#2469. Convert the Temperature
You are given a non-negative floating point number rounded to two decimal places celsius, that denotes the temperature in Celsius.
You should convert Celsius into Kelvin and Fahrenheit and return it as an array ans = [kelvin, fahrenheit].
Return the array ans. Answers within 10-5 of the actual answer will be accepted.

Note that:
Kelvin = Celsius + 273.15
Fahrenheit = Celsius * 1.80 + 32.00

Example 1:
Input: celsius = 36.50
Output: [309.65000,97.70000]
Explanation: Temperature at 36.50 Celsius converted in Kelvin is 309.65 and converted in Fahrenheit is 97.70.

Example 2:
Input: celsius = 122.11
Output: [395.26000,251.79800]
Explanation: Temperature at 122.11 Celsius converted in Kelvin is 395.26 and converted in Fahrenheit is 251.798.

Constraints:
0 <= celsius <= 1000

rus
#2469. Конвертация температуры
Дано неотрицательное число с плавающей точкой, округленное до двух знаков после запятой,
которое обозначает температуру по шкале Цельсия. Вы должны преобразовать градусы Цельсия
в градусы Кельвина и Фаренгейта и вернуть в виде массива ans = [кельвин, фаренгейт].

Верните массив ans. Принимаются ответы в пределах 10-5 от фактического ответа (это означает,
что результат должен быть точен до 5 знаков после запятой. То есть, ваше решение может отличаться
от точного ответа не более чем на 0.00001, и такая погрешность будет считаться допустимой).

Обратите внимание, что:
Кельвин = Цельсий + 273,15
Фаренгейт = Цельсий * 1,80 + 32,00

Пример 1:
Ввод: Цельсий = 36,50
Вывод: [309,65000,97,70000]
Пояснение: Температура при 36,50 Цельсия, преобразованная в Кельвины, составляет 309,65, а преобразованная в Фаренгейты — 97,70.

Пример 2:
Ввод: Цельсий = 122,11
Вывод: [395,26000,251,79800]
Пояснение: Температура при 122,11 Цельсия, преобразованная в Кельвины, составляет 395,26, а преобразованная в Фаренгейты — 251,798.

Ограничения:
0 <= по Цельсию <= 1000
*/

func ConvertTemperature(celsius float64) []float64 {

	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00

	return []float64{kelvin, fahrenheit}

}
