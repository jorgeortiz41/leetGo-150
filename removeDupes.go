package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 0 //unique elements

	if len(nums) == 1 {
		return k + 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[k] {
			continue
		} else {
			nums[k+1] = nums[i]
			k++
		}
	}

	return k + 1
}

func TestRemoveDuplicates() {
	// Test case 1
	nums1 := []int{1, 1, 2}
	k := removeDuplicates(nums1)
	fmt.Println(nums1)
	fmt.Println(k)

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k = removeDuplicates(nums2)
	fmt.Println(nums2)
	fmt.Println(k)
}
