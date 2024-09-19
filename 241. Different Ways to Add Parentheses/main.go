/*
Given a string expression of numbers and operators, return all possible results
 from computing all the different possible ways to group numbers and operators. You may return the answer in any order.

The test cases are generated such that the output values fit in a 32-bit integer and the number of different results does not exceed 104.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}

func diffWaysToCompute(expression string) []int {
	var nums []int
	var ops []rune
	i := 0
	for i < len(expression) {
		if isDigit(rune(expression[i])) {
			num := 0
			for i < len(expression) && isDigit(rune(expression[i])) {
				num = num*10 + int(expression[i]-'0')
				i++
			}
			nums = append(nums, num)
		} else {
			ops = append(ops, rune(expression[i]))
			i++
		}
	}

	n := len(nums)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
	}

	for j := 0; j < n; j++ {
		dp[j][j] = []int{nums[j]}
	}

	for length := 2; length <= n; length++ {
		for l := 0; l <= n-length; l++ {
			r := l + length - 1
			for k := l; k < r; k++ {
				leftResults := dp[l][k]
				rightResults := dp[k+1][r]
				operator := ops[k]

				for _, left := range leftResults {
					for _, right := range rightResults {
						if operator == '+' {
							dp[l][r] = append(dp[l][r], left+right)
						} else if operator == '-' {
							dp[l][r] = append(dp[l][r], left-right)
						} else {
							dp[l][r] = append(dp[l][r], left*right)
						}
					}
				}
			}
		}
	}

	return dp[0][n-1]
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
