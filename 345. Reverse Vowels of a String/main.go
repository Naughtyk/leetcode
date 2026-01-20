/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:

Input: s = "IceCreAm"

Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"

Output: "leotcede"
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	indexes := make([]int, 0)

	for i, char := range s {
		if char == 'a' || char == 'e' || char == 'u' || char == 'o' || char == 'i' ||
			char == 'A' || char == 'E' || char == 'U' || char == 'O' || char == 'I' {
			indexes = append(indexes, i)
		}
	}

	ret := make([]rune, len(s))
	right := len(indexes) - 1
	left := 0
	for i, char := range s {
		if right >= 0 && indexes[left] == i {
			ret[i] = rune(s[indexes[right]])
			right -= 1
			left += 1
		} else {
			ret[i] = char

		}
	}
	return string(ret)
}
