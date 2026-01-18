/*
Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

Example 1:


Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
Example 2:

Input: root = [1]
Output: ["1"]
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	ret := make([]string, 0)
	if root.Left == nil && root.Right == nil {
		ret = append(ret, fmt.Sprintf("%d", root.Val))
	}
	if root.Left != nil {
		for _, val := range binaryTreePaths(root.Left) {
			ret = append(ret, fmt.Sprintf("%d->%s", root.Val, val))
		}

	}
	if root.Right != nil {
		for _, val := range binaryTreePaths(root.Right) {
			ret = append(ret, fmt.Sprintf("%d->%s", root.Val, val))
		}
	}
	return ret
}

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{Val: 1,
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 2},
			Right: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6}}}}))
}
