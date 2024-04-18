package main

import "fmt"

func findMax(nums []int) int {
	maxIndex := 0
	maxSum := nums[0] + 0

	for i := 1; i < len(nums); i++ {
		sum := nums[i] + i
		if sum > maxSum {
			maxSum = sum
			maxIndex = i
		}
	}

	return maxIndex + 1
}

func canJump2(nums []int) int {
	jumps := 0
	i := 0

	for i < len(nums)-1 {
		if i+nums[i] >= len(nums)-1 {
			jumps++
			break
		}
		if i+nums[i] <= len(nums)-1 {
			i += findMax(nums[i+1 : i+nums[i]+1])
			jumps++
		} else {
			jumps++
			break
		}
	}

	return jumps
}

func TestJumpGame2() {
	// Test case 1
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println("nums1:", nums1)
	fmt.Println("canJump1:", canJump2(nums1))
	fmt.Println("expected:", 2)

	// Test case 2
	nums2 := []int{3, 2, 1, 1, 4}
	fmt.Println("nums2:", nums2)
	fmt.Println("canJump2:", canJump2(nums2))
	fmt.Println("expected:", 2)

	// Test case 3
	nums3 := []int{1}
	fmt.Println("nums3:", nums3)
	fmt.Println("canJump3:", canJump2(nums3))
	fmt.Println("expected:", 0)

	// Test case 4
	nums4 := []int{4, 1, 1, 3, 1, 1, 1}
	fmt.Println("nums4:", nums4)
	fmt.Println("canJump4:", canJump2(nums4))
	fmt.Println("expected:", 2)

	// Test case 5
	nums5 := []int{2, 3, 1}
	fmt.Println("nums5:", nums5)
	fmt.Println("canJump5:", canJump2(nums5))
	fmt.Println("expected:", 1)
}
