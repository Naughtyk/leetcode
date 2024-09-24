/*
You are given two arrays with positive integers arr1 and arr2.

A prefix of a positive integer is an integer formed by one or more of its digits,
 starting from its leftmost digit. For example, 123 is a prefix of the integer 12345, while 234 is not.

A common prefix of two integers a and b is an integer c, such that c is a prefix
of both a and b. For example, 5655359 and 56554 have a common prefix 565 while 1223 and 43456 do not have a common prefix.

You need to find the length of the longest common prefix between all pairs of integers
 (x, y) such that x belongs to arr1 and y belongs to arr2.

Return the length of the longest common prefix among all pairs.
If no common prefix exists among them, return 0.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	s := make(map[int]struct{})
	ret := 0
	for _, elem := range arr1 {
		for elem != 0 {
			s[elem] = struct{}{}
			elem /= 10
		}
	}
	for _, elem := range arr2 {
		j := len(strconv.Itoa(elem))
		for elem != 0 {
			if j <= ret {
				break
			}
			_, ok := s[elem]
			if ok {
				ret = max(ret, j)
				break
			}
			j--
			elem /= 10
		}
	}
	return ret
}
