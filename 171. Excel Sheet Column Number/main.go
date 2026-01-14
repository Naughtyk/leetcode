/*
Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


Example 1:

Input: columnTitle = "A"
Output: 1
Example 2:

Input: columnTitle = "AB"
Output: 28
Example 3:

Input: columnTitle = "ZY"
Output: 701
*/

package main

import (
	"fmt"
)

func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	res := x
	for range y - 1 {
		res *= x
	}

	return res
}

func titleToNumber(columnTitle string) int {
	ret := 0
	runes := []rune(columnTitle)

	for i := range len(runes) {
		ret += (int(runes[i]) - 64) * pow(26, len(runes)-i-1)
	}
	return ret
}

func main() {
	fmt.Println(titleToNumber("A"))       // 1
	fmt.Println(titleToNumber("B"))       // 2
	fmt.Println(titleToNumber("Z"))       // 26
	fmt.Println(titleToNumber("AB"))      // 28
	fmt.Println(titleToNumber("ZY"))      // 701
	fmt.Println(titleToNumber("ZYC"))     // 18229
	fmt.Println(titleToNumber("FXSHRXW")) // 2 147 483 647

}
