package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}
		if prefix == strs[i][:len(prefix)] {
			continue
		} else {
			prefix = prefix[:len(prefix)-1]
			i = i - 1 // Check against the same string again
		}
	}
	return prefix
}

func TestlongestCommonPrefix() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	strs3 := []string{"a", "a", "a"}
	strs4 := []string{"pantry", "pants", "pantheon"}

	fmt.Println("OUT: ", longestCommonPrefix(strs1), "Expected: fl")
	fmt.Println("OUT: ", longestCommonPrefix(strs2), "Expected: ")
	fmt.Println("OUT: ", longestCommonPrefix(strs3), "Expected: a")
	fmt.Println("OUT: ", longestCommonPrefix(strs4), "Expected: pant")
}
