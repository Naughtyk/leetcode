/*
You are given an integer array score of size n, where score[i] is
 the score of the ith athlete in a competition. All the scores are
  guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place
 athlete has the highest score, the 2nd place athlete has the 2nd
 highest score, and so on. The placement of each athlete determines their rank:

The 1st place athlete's rank is "Gold Medal".
The 2nd place athlete's rank is "Silver Medal".
The 3rd place athlete's rank is "Bronze Medal".
For the 4th place to the nth place athlete, their rank is their
 placement number (i.e., the xth place athlete's rank is "x").
Return an array answer of size n where answer[i] is the rank of
 the ith athlete.
*/

package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Equal(findRelativeRanks([]int{5, 4, 3, 2, 1}), []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}))
	fmt.Println(Equal(findRelativeRanks([]int{10, 3, 8, 9, 4}), []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}))
}

func findRelativeRanks(score []int) []string {
	ret := make([]string, len(score))
	score_copy := make([]int, len(score))

	copy(score_copy, score)

	sort.Ints(score)
	slices.Reverse(score)
	dict := make(map[int]string) // key is score, value is rank
	for i, v := range score {
		switch i {
		case 0:
			dict[v] = "Gold Medal"
		case 1:
			dict[v] = "Silver Medal"
		case 2:
			dict[v] = "Bronze Medal"
		default:
			dict[v] = strconv.Itoa(i + 1)
		}
	}
	for i, v := range score_copy {
		ret[i] = dict[v]
	}
	return ret
}
