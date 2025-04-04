package easy

/*
eng
#169. Majority Element

Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2

Constraints:
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
*/

/*
rus
#169. Элемент большинства
Дан срез nums размера n, вернуть наиболее часто встречающийся элемент.
Элемент большинства — это элемент, который встречается более ⌊n / 2⌋ раз.
Можно предположить, что элемент большинства всегда существует в срезе.

Пример 1:
Входные данные: nums = [3,2,3]
Выходные данные: 3

Пример 2:
Входные данные: nums = [2,2,1,1,1,2,2]
Выходные данные: 2

Ограничения:
n == длина среза nums
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
*/

/*
Для решения этой задачи использован алгоритм Бойера-Мура (Boyer-Moore Majority Vote Algorithm),
который не сработает, если будет нарушено вот это условие задачи: You may assume that the majority element
always exists in the array (Можно предположить, что элемент большинства всегда существует в срезе). Это означет, что условие предполагает,
что элемент, который встречается чаще остальных в срезе есть и не предполагает ситуации, когда
несколько элементов встречаются одинаково бошльное число раз, напр., []int{1, 1, 2, 2, 3, 3} или
[]int{1, 2, 3, 4}.
*/

func MajorityElement(nums []int) int {
	count := 0
	var majority int

	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if num == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}
