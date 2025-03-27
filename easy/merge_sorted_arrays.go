package easy

/*
eng
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

Example 3:
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

Constraints:

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109

Follow up: Can you come up with an algorithm that runs in O(m + n) time?
*/

/*
rus
Вам даны два целочисленных массива nums1 и nums2, отсортированных в неубывающем порядке, и два целых числа m и n,
представляющих количество элементов в nums1 и nums2 соответственно.
Объедините nums1 и nums2 в один массив, отсортированный в неубывающем порядке.
Окончательный отсортированный массив не должен возвращаться функцией, а вместо этого должен храниться внутри массива nums1.
Чтобы учесть это, nums1 имеет длину m + n, где первые m элементов обозначают элементы, которые должны быть объединены,
а последние n элементов устанавливаются в 0 и должны игнорироваться. nums2 имеет длину n.

Пример 1:
Входные данные: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Выходные данные: [1,2,2,3,5,6]
Пояснение: Массивы, которые мы объединяем, — [1,2,3] и [2,5,6].
Результат слияния — [1,2,2,3,5,6] с подчеркнутыми элементами, взятыми из nums1.

Пример 2:
Входные данные: nums1 = [1], m = 1, nums2 = [], n = 0
Выходные данные: [1]
Пояснение: Массивы, которые мы объединяем, — [1] и [].
Результат слияния — [1].

Пример 3:
Вход: nums1 = [0], m = 0, nums2 = [1], n = 1
Выход: [1]
Пояснение: Массивы, которые мы объединяем, — это [] и [1].
Результат слияния — [1].
Обратите внимание, что поскольку m = 0, в nums1 нет элементов. 0 нужен только для того, чтобы результат слияния поместился в nums1.

Ограничения:

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109

Продолжение: Можете ли вы придумать алгоритм, который выполняется за время O(m + n)?
*/

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	//i — индекс последнего значащего элемента в nums1 (m-1).
	i := m - 1
	//j — индекс последнего элемента в nums2 (n-1).
	j := n - 1
	//k — индекс последней позиции в nums1, куда мы будем вставлять элементы (m+n-1).
	k := m + n - 1

	//Запускаем цикл, пока в nums2 есть элементы (то есть j >= 0).
	for j >= 0 {
		//Если nums1[i] > nums2[j], значит, элемент nums1[i] больше, и мы должны поместить его в конец nums1.
		if i >= 0 && nums1[i] > nums2[j] {
			//Копируем nums1[i] в nums1[k] и уменьшаем i.
			nums1[k] = nums1[i]
			i--
		} else {
			//Иначе, если nums2[j] больше (или i < 0, то есть nums1 закончился), вставляем nums2[j] в nums1[k] и уменьшаем j.
			nums1[k] = nums2[j]
			j--
		}
		//Уменьшаем k, чтобы двигаться к следующей позиции в nums1.
		k--
	}
}
