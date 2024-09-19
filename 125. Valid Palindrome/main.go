/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome(".,"))
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		for (!unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l]))) && l < len(s) {
			l += 1
			if l >= len(s) {
				return true
			}
		}
		for (!unicode.IsLetter(rune(s[r]))) && !unicode.IsNumber(rune(s[r])) && r >= 0 {
			r -= 1
			if r < 0 {
				return true
			}
		}
		if !strings.EqualFold(string(s[l]), string(s[r])) {
			return false
		}

		l += 1
		r -= 1
	}
	return true
}
