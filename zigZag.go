package main

import "fmt"

// 6. ZigZag Conversion
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
// P   A   H   N
// A P L S I I G
// Y   I   R
//
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows.

func convert(s string, numRows int) string {
	strRunes := make([]string, numRows)
	rowIndex := 0
	direction := "down"
	result := ""

	if numRows == 1 {
		return s
	}

	for _, sub := range s {
		strRunes[rowIndex] += string(sub)
		if direction == "down" {
			rowIndex++
			if rowIndex > len(strRunes)-1 {
				direction = "up"
				rowIndex = len(strRunes) - 2
			}
		} else {
			rowIndex--
			if rowIndex < 0 {
				rowIndex = 1
				direction = "down"
			}
		}
	}

	for _, str := range strRunes {
		result += str
	}

	return result
}

func TestConvert() {
	fmt.Println("OUT: ", convert("PAYPALISHIRING", 3), "Expected: PAHNAPLSIIGYIR")
	if convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR" {
		fmt.Println("PASSED!")
	} else {
		fmt.Println("FAILED!")
	}
	fmt.Println("OUT: ", convert("AB", 1), "Expected: AB")
	if convert("AB", 1) == "AB" {
		fmt.Println("PASSED!")
	} else {
		fmt.Println("FAILED!")
	}
}
