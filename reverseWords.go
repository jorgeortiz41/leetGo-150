package main

import (
	"fmt"
	"strings"
)

// 151. Reverse Words in a String
// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
//
// Example 1:
// Input: s = "the sky is blue"
// Output: "blue is sky the"

func reverseWords(s string) string {
	reversed := ""
	s = strings.Trim(s, " ")
	fmt.Println("trimmed str: ", s)
	splitStr := strings.Split(s, " ")

	for i := len(splitStr) - 1; i >= 0; i-- {
		fmt.Println("string[i]: ", splitStr[i], "length: ", len(splitStr[i]))
		strings.Trim(splitStr[i], " ")
		if len(splitStr[i]) != 0 {
			reversed += splitStr[i]
			if i != 0 {
				reversed += " "
			}
		}
	}

	return reversed
}

func TestreverseWords() {
	fmt.Println("OUT: ", reverseWords("the sky is blue"), "Expected: blue is sky the")
	fmt.Println("OUT: ", reverseWords("  hello world  "), "Expected: world hello")
	fmt.Println("OUT: ", reverseWords("a good   example"), "Expected: example good a")
	fmt.Println("OUT: ", reverseWords("  Bob    Loves  Alice   "), "Expected: Alice Loves Bob")
	fmt.Println("OUT: ", reverseWords("Alice does not even like bob"), "Expected: bob like even not does Alice")
}
