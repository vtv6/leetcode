package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	x, _ := calc(head, n)
	return x
}

func calc(head *ListNode, n int) (*ListNode, int) {

	if head == nil {
		return nil, 0
	}

	next, pre := calc(head.Next, n)
	head.Next = next
	if pre == n-1 {
		return head.Next, pre + 1
	}

	return head, pre + 1
}
