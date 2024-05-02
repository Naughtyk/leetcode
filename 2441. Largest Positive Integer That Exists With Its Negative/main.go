/*
Given an integer array nums that does
 not contain any zeros, find the
  largest positive integer k such
   that -k also exists in the array.

Return the positive integer k.
 If there is no such integer,
  return -1.

*/

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findMaxK([]int{648, 674, 610}))
}

func findMaxK(nums []int) int {
	ret := -1
	slices.SortFunc(nums, func(i, j int) int {
		return cmp.Compare(i, j)
	})
	right := -1
	for idx, num := range nums {
		if num > 0 {
			right = idx
			break
		}
	}
	if right == -1 {
		return ret
	}
	left := right - 1

	for right < len(nums) {
		for left > 0 && nums[right] > -nums[left] {
			left -= 1
		}
		if left == -1 {
			return ret
		}
		if nums[right] == -nums[left] {
			ret = nums[right]
		}
		right += 1
	}

	return ret
}
