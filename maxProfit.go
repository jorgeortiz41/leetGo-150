package main

import (
	"fmt"
	"math"
)

// maxProfit returns the maximum profit from buying and selling a stock.
// The input is an array of integers where the ith element is the price of a given stock on day i.
// The function should return the maximum profit that can be achieved from buying on one day and selling on another.
// If no profit can be made, return 0.

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit

}

func TestMaxProfit() {
	// Test case 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("prices1:", prices1)
	fmt.Println("maxProfit1:", maxProfit(prices1))
	fmt.Println("expected:", 5)

	// Test case 2
	prices2 := []int{2, 1, 2, 0, 1}
	fmt.Println("prices2:", prices2)
	fmt.Println("maxProfit2:", maxProfit(prices2))
	fmt.Println("expected:", 1)

	// Test case 3
	prices3 := []int{7, 1, 5, 3, 6, 4, 9}
	fmt.Println("prices3:", prices3)
	fmt.Println("maxProfit3:", maxProfit(prices3))
	fmt.Println("expected:", 8)

	// Test case 4
	prices4 := []int{7, 1}
	fmt.Println("prices4:", prices4)
	fmt.Println("maxProfit4:", maxProfit(prices4))
	fmt.Println("expected:", 0)

	// Test case 5
	prices5 := []int{7}
	fmt.Println("prices5:", prices5)
	fmt.Println("maxProfit5:", maxProfit(prices5))
	fmt.Println("expected:", 0)
}
