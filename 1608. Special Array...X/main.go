/*
You are given an array nums of non-negative integers. nums is considered special if there exists a number x such that there are
 exactly x numbers in nums that are greater than or equal to x.

Notice that x does not have to be an element in nums.

Return x if the array is special, otherwise, return -1. It can be proven that if nums is special,
 the value for x is unique.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(specialArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(specialArray([]int{3, 5}))
	fmt.Println(specialArray([]int{0, 0}))
}

func specialArray(nums []int) int {
	countMap := map[int]int{}

	for _, v := range nums {
		countMap[v] += 1
	}
	ans := -1

	for i := 1; i <= len(nums); i++ {
		count := 0
		for key, val := range countMap {
			if key >= i {
				count += val
			}
		}
		if count == i {
			ans = i
		}
	}
	return ans
}
