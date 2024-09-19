/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.
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
