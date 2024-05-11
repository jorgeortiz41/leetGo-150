package main

import "fmt"

// 13. Roman to Integer
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol 	 Value
// I 	     1
// V 	     5
// X 	     10
// L 	     50
// C 	     100
// D 	     500
// M 	     1000
//
// 'I' can be placed before 'V' (5) and 'X' (10) to make 4 and 9.
// 'X' can be placed before 'L' (50) and 'C' (100) to make 40 and 90.
// 'C' can be placed before 'D' (500) and 'M' (1000) to make 400 and 900.
//
// Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

func romanToInt(s string) int {
	letters := []rune(s)
	romans := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	prev := romans[letters[len(letters)-1]] // Last letter
	result := prev

	for i := len(letters) - 2; i >= 0; i-- {
		current := romans[letters[i]]
		if current < prev {
			result -= current
			fmt.Println("current < prev, i: ", i, "result: ", result)
		} else {
			result += current
			fmt.Println("else, i: ", i, "result: ", result)
		}
		prev = current
	}
	return result
}

func TestRomanToInt() {
	fmt.Println("OUT: ", romanToInt("III"), "Expected: 3")
	fmt.Println("OUT: ", romanToInt("IV"), "Expected: 4")
	fmt.Println("OUT: ", romanToInt("IX"), "Expected: 9")
	fmt.Println("OUT: ", romanToInt("LVIII"), "Expected: 58")
	fmt.Println("OUT: ", romanToInt("MCMXCIV"), "Expected: 1994")
}
