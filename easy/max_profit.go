package easy

/*
eng
#121. Best Time To Buy And Sell Stock
You are given an array prices where prices[i] is the price of a given stock on the i^th day.
You want to maximize your profit by choosing a single day to buy one stock and choosing
a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:
1 <= prices.length <= 105
0 <= prices[i] <= 104
*/

/*
rus
#121. Лучшее время для покупки и продажи акций
Дан срез цен, где цены[i] — это цена данной акции в i-й день («Купим акцию в 1-й день, а продадим в 4-й день»).
Вы хотите максимизировать свою прибыль, выбрав один день для покупки одной акции и
выбрав другой день в будущем для продажи этой акции.
Верните максимальную прибыль, которую вы можете получить от этой транзакции.
Если вы не можете получить никакой прибыли, верните 0.

Пример 1:
Вход: цены = [7,1,5,3,6,4]
Выход: 5
Объяснение: купить на 2-й день (цена = 1) и продать на 5-й день (цена = 6), прибыль = 6-1 = 5.
Обратите внимание, что покупка на 2-й день и продажа на 1-й день не допускаются, поскольку вы должны купить, прежде чем продать.

Пример 2:
Вход: цены = [7,6,4,3,1]
Выход: 0
Пояснение: в этом случае транзакции не производятся, а максимальная прибыль = 0.

Ограничения:
1 <= длина среза цен <= 105
0 <= цена <= 104
*/

func MaxProfit(prices []int) int {
	//Самая низкая цена, по которой мы могли бы купить акцию на данный момент.
	minPrice := prices[0]
	maxProfit := 0

	//Идем по срезу слева направо
	for i := 1; i < len(prices); i++ {
		//Если текущая цена ниже minPrice, обновляем minPrice.
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			//В противном случае считаем прибыль, если бы купили по minPrice, а продали по текущей цене — и сравниваем с maxProfit.
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
