/*
You are given an integer array nums and an integer k.

An array is considered balanced if the value of its maximum element is at most k times the minimum element.

You may remove any number of elements from nums​​​​​​​ without making it empty.

Return the minimum number of elements to remove so that the remaining array is balanced.

Note: An array of size 1 is considered balanced as its maximum and minimum are equal, and the condition always holds true.
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minRemoval([]int{1, 6, 2, 9}, 3)) // 2
}

func minRemoval(nums []int, k int) int {
	n := len(nums)

	if n <= 1 {
		return 0
	}

	slices.Sort(nums)

	left, right := 0, 1
	ret := n
	for right < n {
		if nums[left]*k >= nums[right] {
			ret = min(ret, n-(right-left+1))
			right++
		} else {
			left++
		}
	}
	return ret
}
