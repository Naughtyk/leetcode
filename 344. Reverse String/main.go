/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.
*/

package main

import (
	"fmt"
)

func main() {
	ret := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(ret)
	fmt.Println(ret)
}
func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
