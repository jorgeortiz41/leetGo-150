package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 274. H-Index (Medium) https://leetcode.com/problems/h-index/
// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return compute the researcher's h-index.
// A scientist has an index h if h of their n papers have at least h citations each, and the other n - h papers have no more than h citations each.
// If there are several possible values for h, the maximum one is taken as the h-index.
//
// Example 1:
// Input: citations = [3,0,6,1,5]
// Output: 3

func hIndex(citations []int) int {
	hIndex := len(citations)
	hits := 0
	i := 0

	for {
		// fmt.Println("hIndex:", hIndex, "i:", i, "hits:", hits, "value:", citations[i])
		if hIndex == 0 {
			return 0
		}
		// Edge case for when h == len(citations)
		if hIndex == len(citations) && citations[i] < hIndex {
			hIndex--
			hits = 0
			i = 0
			continue
		}
		//if citations is >= to current h then add a hit and go next, if not, just go to next
		if citations[i] >= hIndex {
			hits++
			i++
		} else {
			i++
		}
		//if hits == hIndex then return that
		if hits == hIndex {
			return hIndex
		}
		//if reached the end
		if i == len(citations) {
			// fmt.Println("reached end")
			i = 0
			hits = 0
			hIndex--
		}
	}
}

func hIndex2(citations []int) int {
	nums := quickSort(citations)

	h := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= h {
			break
		}

		h++
	}

	return h
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot, ptr := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[ptr] >= nums[i] {
			nums[pivot], nums[i] = nums[i], nums[pivot]
			if ptr == pivot {
				ptr = i
			}
			pivot++
		}
	}
	nums[ptr], nums[pivot] = nums[pivot], nums[ptr]

	quickSort(nums[:pivot])
	quickSort(nums[pivot+1:])

	return nums
}

// GenerateRandomInput generates a random slice of integers for testing hIndex function
func GenerateRandomInput(size int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator with current time
	citations := make([]int, size)
	for i := range citations {
		citations[i] = rand.Intn(size) // Generate random integers between 0 and size
	}
	return citations
}

func TestHIndex() {
	// Test inputs
	citations1 := []int{3, 0, 6, 1, 5}
	citations2 := []int{1, 3, 1}
	citations3 := []int{0}
	citations4 := GenerateRandomInput(1000)
	citations5 := GenerateRandomInput(5000)

	// Channel to communicate results
	ch := make(chan int, 2)

	// TEST 1
	start1 := time.Now()
	go func() {
		h := hIndex(citations1)
		ch <- h
	}()

	go func() {
		h := hIndex2(citations1)
		ch <- h
	}()
	result1 := <-ch
	elapsed1 := time.Since(start1)
	result2 := <-ch
	elapsed2 := time.Since(start1)

	// TEST 2
	start2 := time.Now()
	go func() {
		h := hIndex(citations2)
		ch <- h
	}()
	go func() {
		h := hIndex2(citations2)
		ch <- h
	}()
	result3 := <-ch
	elapsed3 := time.Since(start2)
	result4 := <-ch
	elapsed4 := time.Since(start2)

	// TEST 3
	start3 := time.Now()
	go func() {
		h := hIndex(citations3)
		ch <- h
	}()
	go func() {
		h := hIndex2(citations3)
		ch <- h
	}()
	result5 := <-ch
	elapsed5 := time.Since(start3)
	result6 := <-ch
	elapsed6 := time.Since(start3)

	// TEST 4
	start4 := time.Now()
	go func() {
		h := hIndex(citations4)
		ch <- h
	}()
	go func() {
		h := hIndex2(citations4)
		ch <- h
	}()
	result7 := <-ch
	elapsed7 := time.Since(start4)
	result8 := <-ch
	elapsed8 := time.Since(start4)

	// TEST 5
	start5 := time.Now()
	go func() {
		h := hIndex(citations5)
		ch <- h
	}()
	go func() {
		h := hIndex2(citations5)
		ch <- h
	}()
	result9 := <-ch
	elapsed9 := time.Since(start5)
	result10 := <-ch
	elapsed10 := time.Since(start5)

	// Print results
	fmt.Println("TestHIndex")
	fmt.Println("TEST 1: ")
	fmt.Println("hIndex: ", "Result 1:", result1, "Elapsed time:", elapsed1)
	fmt.Println("hIndex2: ", "Result 2:", result2, "Elapsed time:", elapsed2)
	fmt.Println("TEST 2: ")
	fmt.Println("hIndex: ", "Result 3:", result3, "Elapsed time:", elapsed3)
	fmt.Println("hIndex2: ", "Result 4:", result4, "Elapsed time:", elapsed4)
	fmt.Println("TEST 3: ")
	fmt.Println("hIndex: ", "Result 5:", result5, "Elapsed time:", elapsed5)
	fmt.Println("hIndex2: ", "Result 6:", result6, "Elapsed time:", elapsed6)
	fmt.Println("TEST 4: ")
	fmt.Println("hIndex: ", "Result 7:", result7, "Elapsed time:", elapsed7)
	fmt.Println("hIndex2: ", "Result 8:", result8, "Elapsed time:", elapsed8)
	fmt.Println("TEST 5: ")
	fmt.Println("hIndex: ", "Result 9:", result9, "Elapsed time:", elapsed9)
	fmt.Println("hIndex2: ", "Result 10:", result10, "Elapsed time:", elapsed10)
}
