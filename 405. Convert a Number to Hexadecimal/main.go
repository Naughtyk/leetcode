/*
Given a 32-bit integer num, return a string representing its hexadecimal representation. For negative integers, two’s complement method is used.

All the letters in the answer string should be lowercase characters, and there should not be any leading zeros in the answer except for the zero itself.

Note: You are not allowed to use any built-in library method to directly solve this problem.

Example 1:

Input: num = 26
Output: "1a"
Example 2:

Input: num = -1
Output: "ffffffff"
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(toHex(-1)) // ffffffff
}

func toHex(num int) string {
	if num < 0 {
		num += 1 << 32
	}
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	res := ""
	for num > 0 {
		x := num % 16
		res = string(hex[x]) + res
		num /= 16
	}
	return res
}
