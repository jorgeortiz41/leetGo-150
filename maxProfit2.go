package main

import (
	"fmt"
)

// 122. Best Time to Buy and Sell Stock II
// Say you have an array prices for which the ith element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete as many transactions as you like
// (i.e., buy one and sell one share of the stock multiple times).

// Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	gainLoss := []int{}
	maxprofit := 0

	for i := 0; i < len(prices)-1; i++ {
		gainLoss = append(gainLoss, prices[i+1]-prices[i])
	}

	for _, gl := range gainLoss {
		if gl > 0 {
			maxprofit += gl
		}
	}

	return maxprofit
}

func TestMaxProfit2() {
	// Test case 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("prices1:", prices1)
	fmt.Println("maxProfit1:", maxProfit2(prices1))
	fmt.Println("expected:", 7)
}
