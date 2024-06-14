/*
You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.

Return the minimum number of moves to make every value in nums unique.

The test cases are generated so that the answer fits in a 32-bit integer.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	ret := 0
	min := nums[0] + 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			ret += min - nums[i]
		} else {
			min = nums[i]
		}
		min += 1
	}
	return ret
}

// 1 1 2 2 3 7
