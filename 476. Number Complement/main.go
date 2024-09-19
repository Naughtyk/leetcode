/*
The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.

For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer num, return its complement.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(7))
	fmt.Println(findComplement(10))
	fmt.Println(findComplement(20))
	fmt.Println(findComplement(32))
}

func findComplement(num int) int {
	s := strconv.FormatInt(int64(num), 2)
	compl := make([]string, len(s))
	for _, ch := range s {
		if ch == '0' {
			compl = append(compl, "1")
		} else {
			compl = append(compl, "0")
		}
	}
	ret, _ := strconv.ParseInt(strings.Join(compl, ""), 2, 64)
	return int(ret)
}
