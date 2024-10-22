/*
Design a stack that supports increment operations on its elements.

Implement the CustomStack class:

CustomStack(int maxSize) Initializes the object with maxSize which is the maximum number of elements in the stack.
void push(int x) Adds x to the top of the stack if the stack has not reached the maxSize.
int pop() Pops and returns the top of the stack or -1 if the stack is empty.
void inc(int k, int val) Increments the bottom k elements of the stack by val. If there are less than k elements in the stack, increment all the elements in the stack.

*/

package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var levels []int64

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levels = make([]int64, 0)
	recurs(root, 0)
	sort.Slice(levels, func(i, j int) bool { return levels[i] > levels[j] })
	if k > len(levels) {
		return -1
	}
	return int64(levels[k-1])
}

func recurs(root *TreeNode, level int) {
	if root != nil {
		if len(levels) < level+1 {
			levels = append(levels, 0)
		}
		levels[level] += int64(root.Val)
		recurs(root.Left, level+1)
		recurs(root.Right, level+1)
	}
}

func main() {
	fmt.Println(kthLargestLevelSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 1))
}
