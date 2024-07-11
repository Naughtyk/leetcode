/*
You are given a string s that consists of lower case English letters and brackets.

Reverse the strings in each pair of matching parentheses, starting from the innermost one.

Your result should not contain any brackets.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
}

func reverseParentheses(s string) string {
	stack := []string{}

	for _, c := range s {
		c := string(c)

		if c != ")" {
			stack = append(stack, c)
		} else {

			curr := []string{}

			for 0 < len(stack) && stack[len(stack)-1] != "(" {
				curr = append(curr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, curr...)
		}
	}
	return strings.Join(stack[:], "")
}
