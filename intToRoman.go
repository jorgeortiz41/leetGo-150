package main

import (
	"fmt"
	"math"
	"strconv"
)

// 12. Integer to Roman
// Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
// Example 1:
// Input: 3
// Output: "III"
// Example 2:
// Input: 4
// Output: "IV"

func IntToRoman(num int) string {
	digitSTR := strconv.Itoa(num)
	power := len(digitSTR) - 1
	romans := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
	subromans := map[int]string{4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	result := ""

	for i, digit := range digitSTR {
		digitInt, _ := strconv.Atoi(string(digit)) // Convert digit to Int
		coeff := math.Pow(10, float64(power-i))    // Calculate coefficient
		tenthVal := float64(digitInt) * coeff      // Multiply digit by coefficient

		if val, ok := subromans[int(tenthVal)]; ok { // Check if value is in subromans
			result += val
		} else if val, ok := romans[int(tenthVal)]; ok { // Check if value is in romans
			result += val
		} else { // Add the values
			if int(tenthVal) > 5 && int(tenthVal) < 10 {
				remainder := int(tenthVal) % 5
				result += romans[5]
				iters := 0

				for iters < remainder {
					result += romans[1]
					iters++
				}
			} else if int(tenthVal) > 50 && int(tenthVal) < 100 {
				remainder := int(tenthVal) % 50
				loops := remainder / 10
				result += romans[50]
				iters := 0

				for iters < loops {
					result += romans[10]
					iters++
				}
			} else if int(tenthVal) > 500 && int(tenthVal) < 1000 {
				remainder := int(tenthVal) % 500
				loops := remainder / 100
				result += romans[500]
				iters := 0

				for iters < loops {
					result += romans[100]
					iters++
				}
			} else {
				iters := 0
				loops := tenthVal / coeff

				for iters < int(loops) {
					result += romans[int(coeff)]
					iters++
				}
			}
		}
	}

	return result
}

func TestIntToRoman() {
	fmt.Println("TestIntToRoman")
	fmt.Println(IntToRoman(3), "Expected: III")
	fmt.Println(IntToRoman(4), "Expected: IV")
	fmt.Println(IntToRoman(8), "Expected: VIII")
	fmt.Println(IntToRoman(20), "Expected: XX")
	fmt.Println(IntToRoman(78), "Expected: LXXVIII")
	fmt.Println(IntToRoman(10), "Expected: X")
	fmt.Println(IntToRoman(40), "Expected: XL")
	fmt.Println(IntToRoman(105), "Expected: CV")
	fmt.Println(IntToRoman(407), "Expected: CDVII")
	fmt.Println(IntToRoman(707), "Expected: DCCVII")
	fmt.Println(IntToRoman(900), "Expected: CM")
	fmt.Println(IntToRoman(3999), "Expected: MMMCMXCIX")
}
