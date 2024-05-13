package main

import "fmt"

// 58. Length of Last Word
// Given a string s consists of some words separated by spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count == 0 {
				continue
			} else {
				break
			}
		}
		count++
	}
	return count
}

func TestlengthOfLastWord() {
	fmt.Println("OUT: ", lengthOfLastWord("Hello World"), "Expected: 5")
}
