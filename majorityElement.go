package main

import "fmt"

// 169. Majority Element (Easy) https://leetcode.com/problems/majority-element/
// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3

func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	result := nums[0]

	if len(nums) <= 2 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]] += 1
		if countMap[nums[i]] > countMap[result] {
			result = nums[i]
		}
	}
	return result
}

func TestMajorityElement() {
	fmt.Println("Running MajorityElement tests")
	fmt.Println("MajorityElement:", majorityElement([]int{3, 2, 3}))
	fmt.Println("MajorityElement:", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
