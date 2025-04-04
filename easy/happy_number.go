package easy

/*
eng
#202. Happy Number

Write an algorithm to determine if a number n is happy.
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:

Input: n = 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

Example 2:
Input: n = 2
Output: false

Constraints:
1 <= n <= 231 - 1
*/

/*
rus
#202. Счастливое число

Напишите алгоритм для определения, является ли число n счастливым.
Счастливое число — это число, определяемое следующим образом:
Начиная с любого положительного целого числа, замените число суммой квадратов его цифр.
Повторяйте процесс, пока число не станет равным 1 (где оно и останется), или пока оно не зациклится бесконечно в цикле, который не включает 1.
Те числа, для которых этот процесс заканчивается на 1, являются счастливыми.
Верните true, если n является счастливым числом, и false, если нет.

Пример 1:

Вход: n = 19
Выход: true
Объяснение:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

Пример 2:
Вход: n = 2
Выход: false

Ограничения:
1 <= n <= 231 - 1
*/

func IsHappy(n int) bool {
	//Создаем map, чтобы отслеживать числа, которые мы уже видели в процессе вычислений. Это поможет обнаружить циклы.
	seen := make(map[int]bool)

	//Запускаем цикл, который будет выполняться до тех пор, пока n не станет равным 1
	//(что означает, что число счастливо) или пока n не появится в seen (что означает, что мы попали в цикл).
	for n != 1 && !seen[n] {
		//Добавляем текущее значение n в map, помечая его как уже встречавшееся.
		seen[n] = true
		//Вычисляем сумму квадратов цифр числа n с помощью вспомогательной функции sumOfSquares и присваиваем результат n.
		n = sumOfSquares(n)
	}
	//Если цикл завершился потому, что n == 1, возвращаем true (число счастливо).
	//Если цикл прервался из-за обнаруженного повторения, возвращаем false.
	return n == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

/*
Алгоритм Floyd’s Cycle Detection Algorithm (или алгоритм зайца и черепахи) —
это классический метод обнаружения циклов в последовательностях.
Он используется, например, в задачах про нахождение цикла в связном списке и определение счастливого числа.
Как он работает?
Вместо хранения всех ранее встреченных чисел (как в решении с map) используются два указателя:
Медленный (slow) — двигается по последовательности шаг за шагом.
Быстрый (fast) — двигается вдвое быстрее (два шага за итерацию).
Если число счастливо, то fast рано или поздно достигнет 1.
Если число несчастливо, то fast и slow попадут в один и тот же цикл.
*/

func IsHappyFloydsCycle(n int) bool {
	slow, fast := n, sumOfSquares(n)

	for fast != 1 && slow != fast {
		slow = sumOfSquares(slow)               // Двигаем slow на 1 шаг
		fast = sumOfSquares(sumOfSquares(fast)) // Двигаем fast на 2 шага
	}

	return fast == 1
}
