package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	tail := head

	for list1 != nil && list2 != nil {
		var newNode *ListNode
		if list1.Val > list2.Val {
			newNode = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			newNode = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		if head == nil {
			head = newNode
		}
		if tail != nil {
			tail.Next = newNode
		}
		tail = newNode
	}

	if list1 == nil {
		if head != nil {
			tail.Next = list2
		} else {
			head = list2
		}

	}
	if list2 == nil {
		if head != nil {
			tail.Next = list1
		} else {
			head = list1
		}
	}
	return head
}
