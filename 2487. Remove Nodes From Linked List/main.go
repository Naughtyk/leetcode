/*
You are given the head of a linked list.

Remove every node which has a node with a greater value anywhere to the right side of it.

Return the head of the modified linked list.
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 5, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 13, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 8,
					}}}}}
	head = removeNodes(head)
	fmt.Println(head)
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head = reverse(head)

	head_ := head

	for head_.Next != nil {
		if head_.Next.Val < head_.Val {
			head_.Next = head_.Next.Next
		} else {
			head_ = head_.Next
		}
	}

	head = reverse(head)

	return head
}

func reverse(node *ListNode) *ListNode {
	prev := ListNode{}.Next

	curr := node

	nxt := node.Next

	for curr.Next != nil {
		curr.Next = prev
		prev = curr
		curr = nxt
		nxt = curr.Next
	}
	curr.Next = prev
	return curr
}
