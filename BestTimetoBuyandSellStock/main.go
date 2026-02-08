package main

import "fmt"

// maxProfit returns the maximum profit from a single buy and sell.
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

func main() {
	tests := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{1, 2, 3, 4, 5},
	}

	for _, t := range tests {
		fmt.Printf("Prices: %v -> Max Profit: %d\n", t, maxProfit(t))
	}
}
