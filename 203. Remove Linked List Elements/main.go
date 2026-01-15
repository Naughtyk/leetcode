/*
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		if head.Val == val {
			return nil
		}
		return head
	}

	if head.Val == val {
		return removeElements(head.Next, val)
	}

	return &ListNode{Val: head.Val, Next: removeElements(head.Next, val)}
}

func main() {
	fmt.Println(removeElements(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}, 2))
}
