/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	var maximum int
	var currentSum int

	i := 0
	for i < k {
		currentSum += nums[i]
		i++
	}
	maximum = currentSum
	for i < len(nums) {
		currentSum += nums[i] - nums[i-k]
		maximum = max(currentSum, maximum)
		i++
	}

	return float64(maximum) / float64(k)
}
