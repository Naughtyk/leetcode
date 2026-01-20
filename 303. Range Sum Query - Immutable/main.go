/*
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/

package main

import (
	"fmt"
)

func main() {
	arr := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(arr.SumRange(0, 2))
	fmt.Println(arr.SumRange(2, 5))
	fmt.Println(arr.SumRange(0, 5))
}

type NumArray struct {
	CumulativeSum []int
}

func Constructor(nums []int) NumArray {
	cum := make([]int, len(nums)+1)
	sum := 0
	for i, num := range nums {
		sum += num
		cum[i+1] = sum
	}
	return NumArray{CumulativeSum: cum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.CumulativeSum[right+1] - this.CumulativeSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
