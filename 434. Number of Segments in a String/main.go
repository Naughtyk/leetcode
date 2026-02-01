/*
Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.

Example 1:

Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
Example 2:

Input: s = "Hello"
Output: 1
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSegments("abc a"))
}

func countSegments(s string) int {
	ret := 0
	cntr := 0
	for _, char := range s {
		cntr++
		if char == ' ' {
			if cntr != 1 {
				ret++
			}
			cntr = 0
		}
	}
	if cntr > 0 {
		ret++
	}
	return ret
}
