/*
Given two strings s1 and s2, return true if s2 contains a
permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.
*/

package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("prosperity", "properties"))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := make([]int, 26)
	s2Count := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if match(s1Count, s2Count) {
			return true
		}
		s2Count[s2[i]-'a']--
		s2Count[s2[i+len(s1)]-'a']++
	}

	return match(s1Count, s2Count)
}

func match(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
