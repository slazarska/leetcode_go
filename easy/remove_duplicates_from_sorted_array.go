package easy

/*
eng
#26. Remove Duplicates From Sorted Array

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place
such that each unique element appears only once. The relative order of the elements should be kept the same.
Then return the number of unique elements in nums.
Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the unique elements in the order
they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.

Example 1:
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:
1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.
*/

/*
rus
#26. Удалить дубликаты из отсортированного массива

Дан целочисленный массив nums, отсортированный в неубывающем порядке, удалить дубликаты на месте так,
чтобы каждый уникальный элемент появлялся только один раз. Относительный порядок элементов должен быть сохранен.
Затем вернуть количество уникальных элементов в nums.
Считайте, что количество уникальных элементов nums равно k, чтобы решить задачу, вам нужно сделать следующее:
Изменить массив nums так, чтобы первые k элементов nums содержали уникальные элементы в том порядке,
в котором они изначально присутствовали в nums. Остальные элементы nums не важны, как и размер nums.
Вернуть k.

Пример 1:
Входные данные: nums = [1,1,2]
Выходные данные: 2, nums = [1,2,_]
Объяснение: Ваша функция должна возвращать k = 2, при этом первые два элемента nums будут равны 1 и 2 соответственно.
Неважно, что вы оставляете после возвращаемого k (следовательно, это подчеркивания).

Пример 2:
Входные данные: nums = [0,0,1,1,1,2,2,3,3,4]
Выходные данные: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Пояснение: Ваша функция должна возвращать k = 5, причем первые пять элементов nums будут 0, 1, 2, 3 и 4 соответственно.
Неважно, что вы оставляете после возвращаемого k (следовательно, это подчеркивания).

Ограничения:
1 <= длина nums <= 3 * 104
-100 <= nums[i] <= 100
nums сортируется в неубывающем порядке.
*/

func RemoveDuplicates(nums []int) int {
	//Проверяем, что массив не пустой, иначе сразу возвращаем ноль
	if len(nums) == 0 {
		return 0
	}

	//Создаем переменную k, которая будет указывать на следующую позицию для записи уникального элемента.
	//Начинаем с 1, так как первый элемент уникален по умолчанию.
	k := 1

	//Идем по массиву, начиная со второго элемента.
	for i := 1; i < len(nums); i++ {
		///Проверяем, отличается ли текущий элемент от предыдущего. Если да, значит, нашли новый уникальный элемент.
		if nums[i] != nums[i-1] {
			//Записываем найденный уникальный элемент на позицию k.
			nums[k] = nums[i]
			//Увеличиваем счетчик k
			k++
		}
	}

	return k
}
