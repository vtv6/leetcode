package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	//  not enough k nodes
	count := 0
	tail := head
	for count < k && tail != nil {
		count++
		tail = tail.Next
	}

	if count < k {
		return head
	}

	temp := reverseKGroup(tail, k)
	pointer := head.Next
	prev := head

	for pointer != tail {
		next := pointer.Next
		pointer.Next = prev
		prev = pointer
		pointer = next
	}
	head.Next = temp

	return prev
}
