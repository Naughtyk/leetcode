/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must appear as many times as it shows in both arrays
and you may return the result in any order.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(intersect([]int{1, 2, 3, 4, 5}, []int{1, 2, 3}))
}

func intersect(nums1 []int, nums2 []int) []int {
	d1 := make(map[int]int)
	d2 := make(map[int]int)
	for _, val := range nums1 {
		d1[val] += 1
	}
	for _, val := range nums2 {
		d2[val] += 1
	}
	ret := []int{}
	for k, v1 := range d1 {
		v2, ok := d2[k]
		if ok {
			for range min(v1, v2) {
				ret = append(ret, k)
			}
		}
	}
	return ret
}
