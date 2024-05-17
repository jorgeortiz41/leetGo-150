package main

import (
	"fmt"
	"reflect"
)

// 68. Text Justification
// Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
// Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
// For the last line of text, it should be left justified and no extra space is inserted between words.
// Note:
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.
//

func fullJustify(words []string, maxWidth int) []string {

}

func TestFullJustify() {
	answer1 := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}

	test1 := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)

	fmt.Println("OUT: ", test1, "Expected: ", answer1)
	if reflect.DeepEqual(test1, answer1) {
		fmt.Println("PASSED!")
	} else {
		fmt.Println("FAILED!")
	}
}
