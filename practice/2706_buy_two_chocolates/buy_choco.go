package buy_two_chocolates

/*
https://leetcode.com/problems/buy-two-chocolates/
solution:

	sort -> pick 2 small prices compare with money
		if sum > money -> return money
		if sum <- money -> return money - sum

	tiume comlexity: O(n*logN)


	iterate all prices => find the smallest price & second smallest price
			if sum > money -> return money
			if sum <- money -> return money - sum
	run time : O(n)
*/
func buyChoco(prices []int, money int) int {
	if money == 1 {
		return 1
	}

	smallestPrice := 100
	secondSmallestPrice := 100

	for _, price := range prices {
		if price < smallestPrice {
			secondSmallestPrice = smallestPrice
			smallestPrice = price
		} else if price >= smallestPrice && price < secondSmallestPrice {
			secondSmallestPrice = price
		}
	}

	sum := smallestPrice + secondSmallestPrice
	if sum > money {
		return money
	} else {
		return money - sum
	}
}
