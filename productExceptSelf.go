package main

import "fmt"

// 238. Product of Array Except Self (Medium) https://leetcode.com/problems/product-of-array-except-self/
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	suffix := 1

	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func TestProductExceptSelf() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println("Running ProductExceptSelf tests")
	fmt.Println("original array:", nums1)
	fmt.Println("ProductExceptSelf:", productExceptSelf(nums1), "Expected output: [24, 12, 8, 6]")
}
