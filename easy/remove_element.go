package easy

/*
eng
#27. Remove Element

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
The remaining elements of nums are not important as well as the size of nums.
Return k.

Example 1:
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:
0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
*/

/*
rus
#27. Удалить элемент

Дан целочисленный массив nums и целочисленный val, удалите все вхождения val в nums на месте (это означает,
что изменения вносятся прямо в исходный массив, без использования дополнительной памяти для нового массива.
То есть, мы не создаем новый массив для хранения результата, а изменяем существующий).
Порядок элементов может быть изменен. Затем верните количество элементов в nums, которые не равны val.
Рассмотрите количество элементов в nums, которые не равны val, как k, вам нужно сделать следующее:
Измените массив nums так, чтобы первые k элементов nums содержали элементы, которые не равны val.
Остальные элементы nums не важны, как и размер nums.
Возвращает k.

Пример 1:
Входные данные: nums = [3,2,2,3], val = 3
Выходные данные: 2, nums = [2,2,_,_]
Объяснение: Ваша функция должна возвращать k = 2, причем первые два элемента nums равны 2.
Неважно, что вы оставите после возвращаемого k (следовательно, это подчеркивания).

Пример 2:
Вход: nums = [0,1,2,2,3,0,4,2], val = 2
Выход: 5, nums = [0,1,4,0,3,_,_,_]
Объяснение: Ваша функция должна возвращать k = 5, причем первые пять элементов nums должны содержать 0, 0, 1, 3 и 4.
Обратите внимание, что пять элементов могут быть возвращены в любом порядке.
Неважно, что вы оставите после возвращаемого k (следовательно, это подчеркивания).

Ограничения:
0 <= длина nums <= 100
0 <= nums[i] <= 50
0 <= val <= 100
*/

func RemoveElement(nums []int, val int) int {
	// счетчик количества элементов, не равных val
	k := 0

	for _, num := range nums {
		if num != val {
			//Если текущий элемент num не равен val, то записываем его в nums[k] (то есть в начало массива, начиная с 0-й позиции).
			nums[k] = num
			//увеличиваем счетчик
			k++
		}
	}

	return k
}
