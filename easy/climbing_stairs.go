package easy

/*
eng
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
1 <= n <= 45
*/

/*
rus
#70. Подъем по лестнице

Вы поднимаетесь по лестнице. Чтобы достичь вершины, нужно сделать n шагов.
Каждый раз вы можете подняться на 1 или 2 ступеньки. Сколькими различными способами вы можете подняться на вершину?

Пример 1:
Вход: n = 2
Выход: 2
Пояснение: Есть два способа подняться на вершину.
1. 1 ступенька + 1 ступенька
2. 2 ступеньки

Пример 2:
Вход: n = 3
Выход: 3
Пояснение: Есть три способа подняться на вершину.
1. 1 ступенька + 1 ступенька + 1 ступенька
2. 1 ступенька + 2 ступеньки
3. 2 ступеньки + 1 ступенька

Ограничения:
1 <= n <= 45
*/

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	ways1, ways2 := 2, 1
	var current int

	for i := 3; i <= n; i++ {
		current = ways1 + ways2
		ways2 = ways1
		ways1 = current
	}

	return current
}
