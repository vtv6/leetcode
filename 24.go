package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	y := head.Next
	nextNode := swapPairs(y.Next)
	head.Next = nextNode
	y.Next = head

	return y
}
