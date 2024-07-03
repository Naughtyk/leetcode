/*
You are given an integer array nums.

In one move, you can choose one element of nums and change it to any value.

Return the minimum difference between the largest
 and smallest value of nums after performing at most three moves.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minDifference([]int{1, 5, 0, 10, 14, 120}))
}

func minDifference(nums []int) int {
	sort.Ints(nums)
	if len(nums) <= 4 {
		return 0
	}
	m := nums[len(nums)-4] - nums[0]
	m = min(m, nums[len(nums)-3]-nums[1])
	m = min(m, nums[len(nums)-2]-nums[2])
	m = min(m, nums[len(nums)-1]-nums[3])
	print(m)
	return m
}
