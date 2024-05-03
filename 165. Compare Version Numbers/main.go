/*
Given two version numbers, version1 and version2, compare them.

Version numbers consist of
 one or more revisions joined
  by a dot '.'. Each revision
   consists of digits and may
    contain leading zeros. Every
	 revision contains at least one
	  character. Revisions are 0-indexed
	   from left to right, with the leftmost
	    revision being revision 0, the next
		 revision being revision 1, and so on.
		  For example 2.5.33 and 0.1 are valid version numbers.

To compare version numbers, compare their revisions
 in left-to-right order. Revisions are compared using
  their integer value ignoring any leading zeros.
   This means that revisions 1 and 001 are considered
    equal. If a version number does not specify a
	 revision at an index, then treat the revision
	  as 0. For example, version 1.0 is less than
	   version 1.1 because their revision 0s are the
	    same, but their revision 1s are 0 and 1 respectively, and 0 < 1.

Return the following:

If version1 < version2, return -1.
If version1 > version2, return 1.
Otherwise, return 0.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.0.1", "2.0.0") == -1)
	fmt.Println(compareVersion("1.01", "1.001") == 0)
	fmt.Println(compareVersion("1.0", "1.0.0") == 0)
	fmt.Println(compareVersion("0.1", "1.1") == -1)
	fmt.Println(compareVersion("1.01", "1.011") == -1)
	fmt.Println(compareVersion("1.0.1", "1") == 1)
	fmt.Println(compareVersion("01", "1") == 0)
	fmt.Println(compareVersion("1.2", "1.10") == -1)
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	end1 := 0
	for i, char := range v1 {
		char = cut(char)
		if char != "0" && char != "" {
			end1 = i
		}
	}
	v1 = v1[:end1+1]

	end2 := 0
	for i, char := range v2 {
		char = cut(char)
		if char != "0" && char != "" {
			end2 = i
		}
	}
	v2 = v2[:end2+1]

	m := min(len(v1), len(v2))
	for i := range m {
		if cmp(v1[i], v2[i]) == -1 {
			return -1
		} else if cmp(v1[i], v2[i]) == 1 {
			return 1
		}
	}
	if len(v1) > len(v2) {
		return 1
	}
	if len(v2) > len(v1) {
		return -1
	}
	return 0
}

func cmp(v1 string, v2 string) int {
	v1 = cut(v1)
	v2 = cut(v2)
	for range min(len(v1), len(v2)) {
		v1_int, _ := strconv.Atoi(v1)
		v2_int, _ := strconv.Atoi(v2)
		if v1_int > v2_int {
			return 1
		} else if v1_int < v2_int {
			return -1
		}
	}
	if len(v1) > len(v2) {
		return 1
	}
	if len(v1) < len(v2) {
		return -1
	}
	return 0
}

func cut(v string) string {
	cut := -1
	for i, char := range v {
		if char == '0' {
			cut = i
		} else {
			break
		}
	}
	return v[cut+1:]
}
