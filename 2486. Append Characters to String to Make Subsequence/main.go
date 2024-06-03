/*
You are given two strings s and t consisting
 of only lowercase English letters.

Return the minimum number of characters that
 need to be appended to the end of s so that
  t becomes a subsequence of s.

A subsequence is a string that can be derived
 from another string by deleting some or no
  characters without changing the order of
   the remaining characters.
*/

package main

import "fmt"

func main() {
	fmt.Println(appendCharacters("coaching", "coding"))
}

func appendCharacters(s string, t string) int {
    i := 0
    j := 0
    for i < len(s) && j < len(t){
        if s[i] == t[j]{
            j += 1
        }
        i += 1
    }
    return len(t) - j
}