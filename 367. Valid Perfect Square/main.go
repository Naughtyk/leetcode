/*
Given a positive integer num, return true if num is a perfect square or false otherwise.

A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as sqrt.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPerfectSquare(1457))
}

func isPerfectSquare(num int) bool {
	for i := range num + 1 {
		if i*i == num {
			return true
		}
		if i*i > num {
			return false
		}
	}
	return false
}
