func maxProfit(prices []int) int {
	minprice := 101
	maxprofit := 0

	for _, price := range prices {
		minprice = min(minprice, price)
		maxprofit = max(maxprofit, price - minprice)
	}

	return maxprofit
}
