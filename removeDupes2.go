package main

import "fmt"

// 80. Remove Duplicates from Sorted Array II
// Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
// Do not allocate extra space for another array; you must do this by modifying the input array in-place with O(1) extra memory.

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// Initialize pointers
	index := 2

	// Iterate through the array starting from the third element
	for i := 2; i < len(nums); i++ {
		// If the current element is different from the element at 'index - 2',
		// it means it's a new element or a new occurrence of the current element
		// Move the 'index' pointer forward and update the value at that position
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}

	// Return the length of the new array with duplicates removed
	return index
}

func TestRemoveDuplicates2() {
	// Test case 1

	nums1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println("nums1 before:", nums1)
	k := removeDuplicates2(nums1)
	fmt.Println("nums1 after:", nums1)
	fmt.Println(k)

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println("nums2 before:", nums2)
	k = removeDuplicates2(nums2)
	fmt.Println("nums2 after:", nums2)
	fmt.Println(k)
}
