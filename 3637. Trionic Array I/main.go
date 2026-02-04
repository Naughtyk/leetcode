/*
You are given an integer array nums of length n.

An array is trionic if there exist indices 0 < p < q < n − 1 such that:

nums[0...p] is strictly increasing,
nums[p...q] is strictly decreasing,
nums[q...n − 1] is strictly increasing.
Return true if nums is trionic, otherwise return false.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isTrionic([]int{1, 3, 5, 4, 2, 6}))
	fmt.Println(isTrionic([]int{2, 1, 3}))
}

func isTrionic(nums []int) bool {
	p, q := 0, 0
	for i := range len(nums) - 1 {
		if nums[i+1] == nums[i] {
			return false
		}
		if nums[i+1] < nums[i] {
			p = i
			break
		}
	}
	if p == 0 {
		return false
	}
	for i := p; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			return false
		}
		if nums[i+1] > nums[i] {
			q = i
			break
		}
	}
	if p == q || q == 0 {
		return false
	}
	for i := q; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			return false
		}
		if nums[i+1] < nums[i] {
			return false
		}
	}
	return true
}
