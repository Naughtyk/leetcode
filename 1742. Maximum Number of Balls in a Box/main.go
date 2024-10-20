/*
You are given a string s consisting only of uppercase English letters.

You can apply some operations to this string where, in one operation, you can remove any occurrence of one of the substrings "AB" or "CD" from s.

Return the minimum possible length of the resulting string that you can obtain.

Note that the string concatenates after removing the substring and could produce new "AB" or "CD" substrings.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBalls(5, 15))
}

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)

	for i := lowLimit; i <= highLimit; i++ {
		num := i
		s := 0
		for num != 0 {
			s += num % 10
			num /= 10
		}
		m[s]++
	}

	ret := 0
	for _, v := range m {
		ret = max(ret, v)
	}
	return ret
}
