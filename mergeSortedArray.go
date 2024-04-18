package main

import (
	"fmt"
)

func isEqual(slice1, slice2 []int) bool {
	// Check if the slices have the same length
	if len(slice1) != len(slice2) {
		return false
	}

	// Compare each element of the slices
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	// If all elements are equal, return true
	return true
}

// 88. Merge Sorted Array
// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has a size equal to m + n such that it has enough space to hold additional elements from nums2.
// Example 1:
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]

// Time Complexity = O(N*M)(NOT GOOD)
func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0 // Index for nums2
	elemBounds := m

	if n == 0 {
		return

	} else if m == 0 {
		for i, num := range nums2 {
			nums1[i] = num
		}

	} else {
		// While last element of nums2 is not zero
		for j < n {
			// Iterate through nums1
			for i := 0; i < len(nums1); i++ {
				// Store current element in separate variable for switching
				current := nums1[i]

				// if the current element in nums2 is less than current element in nums1, then switch elements
				if nums2[j] < current && i != elemBounds {
					nums1[i] = nums2[j]
					nums2[j] = current
				}

				// If iterated over relevant elements switch and break loop
				if i == elemBounds {
					nums1[i] = nums2[j]
					nums2[j] = current
					break
				}
			}
			elemBounds++
			j++ // Update nums2 index after iteration is done or broken
		}
	}
}

func MergeTests() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	nums3 := []int{1}
	nums4 := []int{}

	nums5 := []int{0}
	nums6 := []int{1}

	nums7 := []int{-12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	nums8 := []int{-49, -45, -42, -41, -40, -39, -39, -39, -38, -36, -34, -34, -33, -33, -32, -31, -29, -28, -26, -26, -24, -21, -20, -20, -18, -16, -16, -14, -11, -7, -6, -5, -4, -4, -3, -3, -2, -2, -1, 0, 0, 0, 2, 2, 6, 7, 7, 8, 10, 10, 13, 13, 15, 15, 16, 17, 17, 19, 19, 20, 20, 20, 21, 21, 22, 22, 24, 24, 25, 26, 27, 29, 30, 30, 30, 35, 36, 36, 36, 37, 39, 40, 41, 42, 45, 46, 46, 46, 47, 48}

	//test1
	fmt.Println("Test 1:")
	merge(nums1, 3, nums2, 3)
	fmt.Println("Nums1:", nums1)
	passed1 := isEqual(nums1, []int{1, 2, 2, 3, 5, 6})
	if passed1 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Failed")
	}

	//test2
	fmt.Println("Test 2:")
	merge(nums3, 1, nums4, 0)
	fmt.Println("Nums1:", nums3)
	passed2 := isEqual(nums3, []int{1})
	if passed2 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Failed")
	}

	//test3
	fmt.Println("Test 3:")
	merge(nums5, 0, nums6, 1)
	fmt.Println("Nums1:", nums5)
	passed3 := isEqual(nums5, []int{1})
	if passed3 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Failed")
	}

	//test4
	fmt.Println("Test 4:")
	merge(nums7, 1, nums8, 90)
	fmt.Println("Nums1:", nums7)
	passed4 := isEqual(nums7, []int{-49, -45, -42, -41, -40, -39, -39, -39, -38, -36, -34, -34, -33, -33, -32, -31, -29, -28, -26, -26, -24, -21, -20, -20, -18, -16, -16, -14, -12, -11, -7, -6, -5, -4, -4, -3, -3, -2, -2, -1, 0, 0, 0, 2, 2, 6, 7, 7, 8, 10, 10, 13, 13, 15, 15, 16, 17, 17, 19, 19, 20, 20, 20, 21, 21, 22, 22, 24, 24, 25, 26, 27, 29, 30, 30, 30, 35, 36, 36, 36, 37, 39, 40, 41, 42, 45, 46, 46, 46, 47, 48})
	if passed4 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Failed")
	}
}
