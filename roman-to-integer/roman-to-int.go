package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MMXXIII"))
}

func romanToInt(s string) int {
	/* Roman numerals are usually written largest to smallest from left to right.
	   However, the numeral for four is not IIII.
	   Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
	   The same principle applies to the number nine, which is written as IX.
	*/
	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return characterMap[s[0]]
	}

	sum := characterMap[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if characterMap[s[i]] < characterMap[s[i+1]] {
			sum -= characterMap[s[i]] // for unqie roman character
		} else {
			sum += characterMap[s[i]]
		}
	}

	return sum
}
