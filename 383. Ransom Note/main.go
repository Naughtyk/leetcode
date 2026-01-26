/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteSet := make(map[rune]int)
	for _, char := range ransomNote {
		ransomNoteSet[char] += 1
	}

	magazineSet := make(map[rune]int)
	for _, char := range magazine {
		magazineSet[char] += 1
	}

	for key, val := range ransomNoteSet {
		v, ok := magazineSet[key]
		if !ok || v < val {
			return false
		}
	}
	return true

}
