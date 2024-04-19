package main

import (
	"fmt"
)

// 274. H-Index (Medium) https://leetcode.com/problems/h-index/
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return compute the researcher's h-index.
// A scientist has an index h if h of their n papers have at least h citations each, and the other n - h papers have no more than h citations each.
// If there are several possible values for h, the maximum one is taken as the h-index.
//
// Example 1:
// Input: citations = [3,0,6,1,5]
// Output: 3

func hIndex(citations []int) int {
	hIndex := len(citations)
	hits := 0
	i := 0

	fmt.Println(len(citations))
	for true {
		fmt.Println("hIndex:", hIndex, "i:", i, "hits:", hits, "value:", citations[i])
		if hIndex == 0 {
			return 0
		}

		// Edge case for when h == len(citations)
		if hIndex == len(citations) && citations[i] < hIndex {
			hIndex--
			hits = 0
			i = 0
		}

		//if citations is >= to current h then add a hit and go next, if not, just go to next
		if citations[i] >= hIndex {
			hits++
			i++
		} else {
			i++
		}

		//if hits == hIndex then return that
		if hits == hIndex {
			return hIndex
		}

		//if reached the end
		if i == len(citations)-1 {
			fmt.Println("reached end")
			i = 0
			hits = 0
			hIndex--
		}
	}

	return hIndex
}

func TestHIndex() {
	citations1 := []int{3, 0, 6, 1, 5}
	fmt.Println("citation1: ", citations1)
	fmt.Println(hIndex(citations1))
	fmt.Println("Expected: 3")
}
