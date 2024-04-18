package main

import "fmt"

// 27. Remove Element
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The relative order of the elements may be changed.
// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
// It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

func removeElement(nums []int, val int) int {
	k := 0
	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left] == val && nums[right] != val {
			nums[left] = nums[right]
			nums[right] = nums[left]
			left++
			right--
			k++
		} else if nums[left] != val && nums[right] == val {
			left++
			right--
			k++
		} else if nums[left] == val && nums[right] == val {
			right--
		} else {
			left++
			k++
		}

	}

	return k
}

func TestRemoveElement() {
	// Test case 1
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val1 := 2
	k := removeElement(nums1, val1)
	fmt.Println(nums1)
	fmt.Println(k)
}
