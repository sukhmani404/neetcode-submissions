func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxxProfit := 0

	for _, price := range prices[1:] {
		minPrice = min(minPrice, price)
		maxxProfit = max(maxxProfit, price - minPrice)
	}

	return maxxProfit
}
