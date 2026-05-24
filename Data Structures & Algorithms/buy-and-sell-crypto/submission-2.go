func maxProfit(prices []int) int {
    l := 0 //left boundary

    if len(prices) < 2 {
        return 0
    }

    maxP := 0

    for r := 1; r < len(prices); r++ {
        if prices[r] > prices[l] {
            maxP = max(maxP, prices[r] - prices[l])
        }else{
            l = r
        }
    }

    return maxP
}
