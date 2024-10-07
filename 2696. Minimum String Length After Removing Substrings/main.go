/*
You are given a string s consisting only of uppercase English letters.

You can apply some operations to this string where, in one operation, you can remove any occurrence of one of the substrings "AB" or "CD" from s.

Return the minimum possible length of the resulting string that you can obtain.

Note that the string concatenates after removing the substring and could produce new "AB" or "CD" substrings.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(minLength("ABFCACDB"))
	fmt.Println(minLength("ACBBD"))
}

func minLength(s string) int {
	stack := make([]rune, 0)
	n := 0
	for _, char := range s {
		stack = append(stack, char)
		n++
		if n > 1 && ((stack[n-2] == 'A' && stack[n-1] == 'B') || (stack[n-2] == 'C' && stack[n-1] == 'D')) {
			n -= 2
			stack = stack[:n]
		}
	}

	return len(stack)
}
