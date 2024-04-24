package main

import (
	"fmt"
)

// 134. Gas Station
// There are N gas stations along a circular route, where the amount of gas at station i is gas[i].
// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1).
// You begin the journey with an empty tank at one of the gas stations.
// Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.
// Note:
// If there exists a solution, it is guaranteed to be unique.
// Both input arrays are non-empty and have the same length.
// Each element in the input arrays is a non-negative integer.
//
// Example 1:
// Input:
// gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]
// Output: 3

// Ideas:
// 1. O(N^2) for each gas station, try to do a full circle
// 2. O(N) find the gas station that has the most gas compared to cost and try full circle, repeat by decreasing diff
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	start := 0
	currTotalGas := 0
	currTotalCost := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currTotalGas += gas[i]
		currTotalCost += cost[i]
		if gas[i] < cost[i] && currTotalGas < currTotalCost {
			start = i + 1
			currTotalGas = 0
			currTotalCost = 0
		}
	}

	if totalGas < totalCost {
		return -1
	} else {
		return start
	}
}

func TestCanCompleteCircuit() {
	gas1 := []int{1, 2, 3, 4, 5}
	cost1 := []int{3, 4, 5, 1, 2}
	fmt.Println("Running test for canCompleteCircuit")
	fmt.Println("gas stations: ", gas1)
	fmt.Println("gas cost: ", cost1)
	fmt.Println("Output: ", canCompleteCircuit(gas1, cost1), "Expected: 3")

	gas2 := []int{2, 3, 4}
	cost2 := []int{3, 4, 3}
	fmt.Println("gas stations: ", gas2)
	fmt.Println("gas cost: ", cost2)
	fmt.Println("Output: ", canCompleteCircuit(gas2, cost2), "Expected: -1")

	gas3 := []int{5, 1, 2, 3, 4}
	cost3 := []int{4, 4, 1, 5, 1}
	fmt.Println("gas stations: ", gas3)
	fmt.Println("gas cost: ", cost3)
	fmt.Println("Output: ", canCompleteCircuit(gas3, cost3), "Expected: 4")

	gas4 := []int{5, 8, 2, 8}
	cost4 := []int{6, 5, 6, 6}
	fmt.Println("gas stations: ", gas4)
	fmt.Println("gas cost: ", cost4)
	fmt.Println("Output: ", canCompleteCircuit(gas4, cost4), "Expected: 3")

	gas5 := []int{3, 1, 1}
	cost5 := []int{1, 2, 2}
	fmt.Println("gas stations: ", gas5)
	fmt.Println("gas cost: ", cost5)
	fmt.Println("Output: ", canCompleteCircuit(gas5, cost5), "Expected: 0")

	gas6 := []int{6, 1, 4, 3, 5}
	cost6 := []int{3, 8, 2, 4, 2}
	fmt.Println("gas stations: ", gas6)
	fmt.Println("gas cost: ", cost6)
	fmt.Println("Output: ", canCompleteCircuit(gas6, cost6), "Expected: 2")
}
