package best_time_to_buy_and_sell_stock_ii

/*
	problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

	time complexity: O(N), N is the length of prices
*/

func maxProfit(prices []int) int {
	valley, peak := prices[0], prices[0]
	maxProfitValue := 0
	i := 0

	for i < len(prices)-1 {

		// finding the valley first, because we need to buy at the valley and sell at the peak
		for i < len(prices)-1 && prices[i+1] <= prices[i] {
			i++
		}
		valley = prices[i]

		//finding the peak
		for i < len(prices)-1 && prices[i+1] >= prices[i] {
			i++
		}
		peak = prices[i]

		maxProfitValue += peak - valley
	}

	return maxProfitValue
}
