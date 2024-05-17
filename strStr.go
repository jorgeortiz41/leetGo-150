package main

import (
	"fmt"
	"strings"
)

// 28. Implement strStr()
// Given a haystack string and a needle string, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of // haystack.

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func TestStrStr() {
	fmt.Println("OUT: ", strStr("hello", "ll"), "Expected: 2")
	fmt.Println("OUT: ", strStr("aaaaa", "bba"), "Expected: -1")
	fmt.Println("OUT: ", strStr("leetcode", "leet"), "Expected: 0")
	fmt.Println("OUT: ", strStr("leetcode", "code"), "Expected: 4")
}
