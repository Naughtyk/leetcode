/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.


Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) // 5 6
	fmt.Println(findDisappearedNumbers([]int{1, 1}))                   // 2
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 4, 2, 3, 1})) // 5 6 8

}

func findDisappearedNumbers(nums []int) []int {
	nums = append(nums, 0)
	n := len(nums)
	ret := make([]int, 0)

	i := 0
	for i < n {
		val := nums[i]
		if val == -1 || nums[val] == -1 {
			i++
		} else if val == i {
			nums[i] = -1
			i++

		} else {
			nums[i] = nums[val]
			nums[val] = -1
		}

	}

	for i, num := range nums {
		if num != -1 {
			ret = append(ret, i)
		}
	}

	return ret
}
