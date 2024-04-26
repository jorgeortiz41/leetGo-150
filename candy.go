package main

import (
	"fmt"
)

// 135. Candy
// There are N children standing in a line. Each child is assigned a rating value.
// You are giving candies to these children subjected to the following requirements:
// - Each child must have at least one candy.
// - Children with a higher rating get more candies than their neighbors.
// What is the minimum candies you must give?
//
// Example 1:
// Input: [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

func candy(ratings []int) int {
	total_candy := 0
	candies := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		candies[i] += 1
		if i > 0 {
			if ratings[i-1] > ratings[i] {
				for j := i; j > 0; j-- {
					if ratings[j-1] <= ratings[j] || candies[j-1] > candies[j] {
						break
					}
					candies[j-1] += 1
				}
			}
		}
		if i < len(ratings)-1 {
			if ratings[i] < ratings[i+1] {
				for j := i; j < len(ratings)-1; j++ {
					if ratings[j] >= ratings[j+1] || candies[j+1] > candies[j] {
						break
					}
					candies[j+1] += 1
				}
			}
		}
	}

	for _, num := range candies {
		total_candy += num
	}
	fmt.Println("candies: ", candies)
	return total_candy
}

func TestCandy() {
	ratings1 := []int{1, 0, 2}
	ratings2 := []int{1, 1, 1}
	ratings3 := []int{1, 2, 3, 4, 5}
	ratings4 := []int{5, 4, 3, 2, 1}
	ratings5 := []int{1, 3, 2, 4, 1, 5}
	ratings6 := []int{1, 2, 2}

	fmt.Println("Running test for candy")
	fmt.Println("OUT(ratings1): ", candy(ratings1), "Expected: 5")
	fmt.Println("OUT(ratings2): ", candy(ratings2), "Expected: 3")
	fmt.Println("OUT(ratings3): ", candy(ratings3), "Expected: 15")
	fmt.Println("OUT(ratings4): ", candy(ratings4), "Expected: 15")
	fmt.Println("OUT(ratings5): ", candy(ratings5), "Expected: 9")
	fmt.Println("OUT(ratings6): ", candy(ratings6), "Expected: 4")
}
