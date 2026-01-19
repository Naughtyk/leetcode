/*
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:

Each letter in pattern maps to exactly one unique word in s.
Each unique word in s maps to exactly one letter in pattern.
No two letters map to the same word, and no two words map to the same letter.


Example 1:

Input: pattern = "abba", s = "dog cat cat dog"

Output: true

Explanation:

The bijection can be established as:

'a' maps to "dog".
'b' maps to "cat".
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"

Output: false

Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"

Output: false
*/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[rune]string)
	m2 := make(map[string]rune)
	sSplitted := strings.Split(s, " ")
	if len(pattern) != len(sSplitted) {
		return false
	}
	for i, char := range pattern {
		word, ok := m[char]
		if !ok {
			if _, ok := m2[sSplitted[i]]; ok {
				return false
			}
			m[char] = sSplitted[i]
			m2[sSplitted[i]] = char
		} else {
			if word != sSplitted[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // true
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // false
	fmt.Println(wordPattern("abba", "dog dog dog dog"))  // false
}
