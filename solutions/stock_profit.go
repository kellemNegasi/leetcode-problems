package solutions

import "math"

func MaxProfit(prices []int) int {
	minPrice := math.MaxInt
	mProfit := 0
	for i := 0; i < len(prices); i++ {

		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		if mProfit < prices[i]-minPrice {
			mProfit = prices[i] - minPrice
		}
	}
	return mProfit
}