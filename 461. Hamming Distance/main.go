/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, return the Hamming distance between them.

Example 1:

Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
Example 2:

Input: x = 3, y = 1
Output: 1
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(1, 4))

}
func hammingDistance(x int, y int) int {
	ret := 0
	for x > 0 && y > 0 {
		if x%2 != y%2 {
			ret++
		}
		x = x >> 1
		y = y >> 1
	}
	for x > 0 {
		if x%2 == 1 {
			ret++
		}
		x = x >> 1
	}
	for y > 0 {
		if y%2 == 1 {
			ret++
		}
		y = y >> 1
	}
	return ret
}
