package main

import (
	"fmt"
)

// 42. Trapping Rain Water
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//
// Example:
//
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
//
// Input: [4,2,0,3,2,5]
// Output: 9

func trap(height []int) int {
	water := 0
	search := false
	lastPeak := 0

	if len(height) <= 2 {
		return 0
	}

	for i := 0; i < len(height)-1; i++ {
		search = true
		j := i + 1
		peak := height[i]
		toFill := 0
		// fmt.Println("i: ", i)
		for search {
			if j == len(height)-1 && height[j] < peak {
				break
			}
			if height[j] < peak {
				toFill += peak - height[j]
				j++
			}
			if height[j] >= peak {
				// fmt.Println("toFill: ", toFill)
				if toFill > 0 {
					lastPeak = j
				}
				water += toFill
				i = j - 1
				break
			}
		}
		// fmt.Println("lastPeak(r-l): ", lastPeak)
	}
	fmt.Println("lastPeak(before l-r): ", lastPeak)
	if lastPeak < len(height)-1 {
		for i := len(height) - 1; i > lastPeak; i-- {
			search = true
			j := i - 1
			peak := height[i]
			toFill := 0

			for search {
				if j == lastPeak && height[j] < peak {
					break
				}
				if height[j] < peak {
					toFill += peak - height[j]
					j--
				}
				if height[j] >= peak {
					water += toFill
					i = j + 1
					break
				}

			}
		}
	}
	return water
}

func TestTrap() {
	height1 := []int{4, 2, 0, 3, 2, 5}
	height2 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height3 := []int{3, 0, 0, 2, 0, 4}
	height4 := []int{1, 2, 3, 4, 5}
	height5 := []int{5, 4, 3, 2, 1}
	height6 := []int{4, 2, 3}
	height7 := []int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3}
	height8 := []int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2}

	fmt.Println("---OUT---(height1): ", trap(height1), "Expected: 9")
	fmt.Println("---OUT---(height2): ", trap(height2), "Expected: 6")
	fmt.Println("---OUT---(height3): ", trap(height3), "Expected: 10")
	fmt.Println("---OUT---(height4): ", trap(height4), "Expected: 0")
	fmt.Println("---OUT---(height5): ", trap(height5), "Expected: 0")
	fmt.Println("---OUT---(height6): ", trap(height6), "Expected: 1")
	fmt.Println("---OUT---(height7): ", trap(height7), "Expected: 83")
	fmt.Println("---OUT---(height8): ", trap(height8), "Expected: 26")
}
