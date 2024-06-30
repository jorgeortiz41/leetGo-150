package main

// 125. Valid Palindrome
// Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool{
	left := 0
	right := len(s)-1

	for left <= right {
		leftrune := rune(s[left])
		rightrune := rune(s[right])
		isLeftAlphaNum := false
		isRightAlphaNum := false

		if (unicode.IsLetter(leftrune) || unicode.IsDigit(leftrune)) {
			isLeftAlphaNum = true
		} else {
			left++
			continue
		}
		if (unicode.IsLetter(rightrune) || unicode.IsDigit(rightrune)) {
			isRightAlphaNum = true
		} else {
			right--
			continue
		}

		if (isLeftAlphaNum && isRightAlphaNum) {
			if (unicode.ToLower(leftrune) == unicode.ToLower(rightrune)) {
				left++
				right--
			} else {
				return false
			}
		}
	}
	return true
}

func TestIsPalindrome(){
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "

	fmt.Println(isPalindrome(s1), "Expected True")
	fmt.Println(isPalindrome(s2), "Expected False")
	fmt.Println(isPalindrome(s3), "Expected True")
}