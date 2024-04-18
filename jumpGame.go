package main

import (
	"fmt"
)

// 55. Jump Game
// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Determine if you are able to reach the last index.
//
// Returns true if you can reach the last index, false otherwise.

func canJump(nums []int) bool {
	lastIndex := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= lastIndex-i {
			lastIndex = i
		}
	}
	return lastIndex == 0
}

func TestJumpGame() {
	// Test case 1
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println("nums1:", nums1)
	fmt.Println("canJump1:", canJump(nums1))
	fmt.Println("expected:", true)

	// Test case 2
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println("nums2:", nums2)
	fmt.Println("canJump2:", canJump(nums2))
	fmt.Println("expected:", false)

	// Test case 3
	nums3 := []int{0}
	fmt.Println("nums3:", nums3)
	fmt.Println("canJump3:", canJump(nums3))
	fmt.Println("expected:", true)

	// Test case 4
	nums4 := []int{3, 0, 8, 2, 0, 0, 1}
	fmt.Println("nums4:", nums4)
	fmt.Println("canJump4:", canJump(nums4))
	fmt.Println("expected:", true)

	// Test case 5
	nums5 := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println("nums5:", nums5)
	fmt.Println("canJump5:", canJump(nums5))
	fmt.Println("expected:", true)
}
