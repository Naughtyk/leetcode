/*
There are n seats and n students in a room. You are given an array seats of length n,
 where seats[i] is the position of the ith seat. You are also given the array students of length n, where students[j] is the position of the jth student.

You may perform the following move any number of times:

Increase or decrease the position of the ith student by 1 (i.e., moving the ith student from position x to x + 1 or x - 1)
Return the minimum number of moves required to move each student to a seat such that no two students are in the same seat.

Note that there may be multiple seats or students in the same position at the beginning.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	moves := 0
	for i := 0; i < len(seats); i++ {
		moves += abs(seats[i] - students[i])
	}
	return moves
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
