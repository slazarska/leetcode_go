package easy

/*
eng
219. Contains Duplicate II

Given an integer array nums and an integer k, return true if there are two distinct indices
i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:
Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:
Input: nums = [1,2,3,1,2,3], k = 2
Output: false

Constraints:
1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/

/*
rus
219. Содержит дубликаты II

Дан срез целых чисел nums и целое число k, вернуть true, если в срезе есть два различных индекса
i и j, такие, что nums[i] == nums[j] и abs(i - j) <= k.

Пример 1:
Вход: nums = [1,2,3,1], k = 3
Выход: true

Пример 2:
Вход: nums = [1,0,1,1], k = 1
Выход: true

Пример 3:
Вход: nums = [1,2,3,1,2,3], k = 2
Выход: false

Ограничения:
1 <= длина nums <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/

func ContainsNearbyDuplicate(nums []int, k int) bool {
	//Создаем пустую мапу, где будем хранить индексы элементов.
	indexMap := make(map[int]int)

	//Проходим по всем элементам среза, где i — текущий индекс, а num — значение элемента.
	for i, num := range nums {
		//Проверяем, встречался ли элемент ранее. Если элемент встречался, то в prevIndex будет храниться его последний индекс.
		if prevIndex, found := indexMap[num]; found {
			//Если разница между текущим и предыдущим индексом меньше или равна k, то возвращаем true.
			if i-prevIndex <= k {
				return true
			}
		}
		//Обновляем мапу, записывая текущий индекс для текущего элемента.
		indexMap[num] = i
	}

	//Если не нашли соответствующих индексов, возвращаем false
	return false
}
