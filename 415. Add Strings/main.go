/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.



Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"
Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"
Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings("11", "123"))
}

func addStrings(num1 string, num2 string) string {
	addition := ""
	idx1, idx2, rem := len(num1)-1, len(num2)-1, 0

	for {
		n1, n2 := 0, 0

		if idx1 < 0 && idx2 < 0 {
			break
		}

		if idx1 >= 0 {
			n1, _ = strconv.Atoi(string(num1[idx1]))
			idx1--
		}

		if idx2 >= 0 {
			n2, _ = strconv.Atoi(string(num2[idx2]))
			idx2--
		}

		sum := n1 + n2 + rem
		digit := sum % 10
		rem = sum / 10
		addition = strconv.Itoa(digit) + addition
	}

	if rem != 0 {
		addition = strconv.Itoa(rem) + addition
	}
	return addition
}
