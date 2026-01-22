/*
Given an array nums, you can perform the following operation any number of times:

Select the adjacent pair with the minimum sum in nums. If multiple such pairs exist, choose the leftmost one.
Replace the pair with their sum.
Return the minimum number of operations needed to make the array non-decreasing.

An array is said to be non-decreasing if each element is greater than or equal to its previous element (if it exists).

Example 1:

Input: nums = [5,2,3,1]

Output: 2

Explanation:

The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
The array nums became non-decreasing in two operations.

Example 2:

Input: nums = [1,2,2]

Output: 0

Explanation:

The array nums is already sorted.
*/

package main

import "fmt"

func main() {
	fmt.Println(minimumPairRemoval([]int{5, 2, 3, 1})) // 2
}

func minimumPairRemoval(nums []int) int {
	counter := 0
	for {
		if len(nums) < 2 {
			return counter
		}
		if check(nums) {
			return counter
		}
		minPair := nums[0] + nums[1]
		idxPair := 0
		for i := range len(nums) - 1 {
			if nums[i]+nums[i+1] < minPair {
				minPair = nums[i] + nums[i+1]
				idxPair = i
			}
		}
		nums[idxPair] = nums[idxPair] + nums[idxPair+1]
		nums = append(nums[:idxPair+1], nums[idxPair+2:]...)
		counter++
	}

}

func check(nums []int) bool {
	for i := range len(nums) - 1 {
		if nums[i+1] < nums[i] {
			return false
		}
	}
	return true
}
