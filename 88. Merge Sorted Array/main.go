/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
 To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be
  merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

package main

import (
	"fmt"
)

func main() {

	sl := []int{1, 2, 3, 0, 0, 0}
	merge(sl, 3, []int{2, 5, 6}, 3)

	fmt.Println(sl)

	sl = []int{4, 5, 6, 0, 0, 0}
	merge(sl, 3, []int{1, 2, 3}, 3)

	fmt.Println(sl)

	sl = []int{0}
	merge(sl, 0, []int{2}, 1)

	fmt.Println(sl)

	sl = []int{1}
	merge(sl, 1, []int{}, 0)

	fmt.Println(sl)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i, j := m-1, n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	if i == -1 {
		for j >= 0 {
			nums1[j] = nums2[j]
			j--
		}
	}
}
