/*
You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.

Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrangeCoins(1)) // 1
	fmt.Println(arrangeCoins(3)) // 2
	fmt.Println(arrangeCoins(5)) // 2
	fmt.Println(arrangeCoins(8)) // 3
}

func arrangeCoins(n int) int {
	s := 0
	i := 1
	for {
		s += i
		if s > n {
			return i - 1
		}
		i++
	}
}
