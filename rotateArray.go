package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	k = k % n // Normalize k to be within the range [0, n)

	if k == 0 {
		return
	}

	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	reverse(0, n-1) // Reverse the entire slice
	reverse(0, k-1) // Reverse the first part up to k
	reverse(k, n-1)
}

func TestRotate() {
	// Test case 1
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("nums1 before:", nums1)
	rotate(nums1, 3)
	fmt.Println("nums1 after:", nums1)

	// Test case 2
	nums2 := []int{-1, -100, 3, 99}
	fmt.Println("nums2 before:", nums2)
	rotate(nums2, 2)
	fmt.Println("nums2 after:", nums2)

	// Test case 3
	// nums3 := []int{1, 2}
	// fmt.Println("nums3 before:", nums3)
	// rotate(nums3, 11)
	// fmt.Println("nums3 after:", nums3)
}
