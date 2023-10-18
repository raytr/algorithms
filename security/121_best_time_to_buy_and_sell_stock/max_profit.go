package best_time_to_buy_and_sell_stock

/*
	problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

   minP = prices[0]
   sellPrice = prices[1]

   let sellPrice pointer traversal the whole of prices
   minP = getMin(minP, sellPrice)
   max = getMax(max, sellPrice - minP)

   return max

   O(n)
*/

func maxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		sellPrice := prices[i]
		minPrice = getMin(minPrice, sellPrice)
		maxProfit = getMax(maxProfit, sellPrice-minPrice)
	}

	return maxProfit
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
