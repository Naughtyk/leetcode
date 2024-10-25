/*
Given the root of a binary tree, replace the value of each node in the tree with the sum of all its cousins' values.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Return the root of the modified tree.

Note that the depth of a node is the number of edges in the path from the root node to it.
*/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var levels []int

func replaceValueInTree(root *TreeNode) *TreeNode {
	levels = make([]int, 0)
	fillLevels(root, 0)
	bfs(root, root.Val, 0)
	return root
}

func lookup(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val
}

func bfs(root *TreeNode, brother, level int) {
	if root == nil {
		return
	}

	root.Val = levels[level] - brother

	childs := lookup(root.Left) + lookup(root.Right)
	bfs(root.Left, childs, level+1)
	bfs(root.Right, childs, level+1)
}

func fillLevels(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(levels) < level+1 {
		levels = append(levels, 0)
	}
	levels[level] += root.Val
	fillLevels(root.Left, level+1)
	fillLevels(root.Right, level+1)
}

func main() {
	fmt.Println(replaceValueInTree(&TreeNode{Val: 5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 10}},
		Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 7}}}))
}
