/*
An ugly number is a positive integer which does not have a prime factor other than 2, 3, and 5.

Given an integer n, return true if n is an ugly number.



Example 1:

Input: n = 6
Output: true
Explanation: 6 = 2 × 3
Example 2:

Input: n = 1
Output: true
Explanation: 1 has no prime factors.
Example 3:

Input: n = 14
Output: false
Explanation: 14 is not ugly since it includes the prime factor 7.
*/

package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for {
		if n%2 == 0 {
			n = n / 2
		} else if n%3 == 0 {
			n = n / 3
		} else if n%5 == 0 {
			n = n / 5
		} else {
			if n == 1 || n == 2 || n == 3 || n == 5 {
				return true
			}
			return false
		}
	}
}

func main() {
	fmt.Println(isUgly(6))  // true
	fmt.Println(isUgly(1))  // true
	fmt.Println(isUgly(14)) // false
	fmt.Println(isUgly(23)) // false

}
