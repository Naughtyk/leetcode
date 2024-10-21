/*
Given a string s, return the maximum number of unique substrings that the given string can be split into.

You can split string s into any list of non-empty substrings,
 where the concatenation of the substrings forms the original
  string. However, you must split the substrings such that all of them are unique.

A substring is a contiguous sequence of characters within a string.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
}

func maxUniqueSplit(s string) int {
	return recurs(make(map[string]bool), s)
}

func recurs(set map[string]bool, s string) int {
	if len(s) == 0 {
		cntr := 0
		for _, v := range set {
			if v {
				cntr++
			}
		}
		return cntr
	}
	ret := 0
	for idx := range s {
		subStr := s[:idx+1]
		val := set[subStr]
		if !val {
			set[subStr] = true
			ret = max(ret, recurs(set, s[idx+1:]))
			set[subStr] = false
		}
	}
	return ret
}
