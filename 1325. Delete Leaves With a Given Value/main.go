/*
Given a binary tree root and an integer target, delete all the leaf nodes with value target.

Note that once you delete a leaf node with value target,
 if its parent node becomes a leaf node and has the value
 target, it should also be deleted (you need to continue doing that until you cannot).

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeLeafNodes(&TreeNode{Val: 5, Left: nil, Right: nil}, 5))
	fmt.Println(removeLeafNodes(&TreeNode{Val: 5, Left: nil, Right: nil}, 4))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
