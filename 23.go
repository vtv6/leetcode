package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, tail *ListNode

	a := []*ListNode{}
	for _, list := range lists {
		if list != nil {
			a = append(a, list)
		}
	}

	lists = a

	for len(lists) > 1 {
		x := lists[0]
		index := 0

		for i, list := range lists {
			if x.Val > list.Val {
				x = list
				index = i
			}
		}

		newNode := &ListNode{Val: x.Val}

		if head == nil {
			head = newNode
		}

		if tail != nil {
			tail.Next = newNode
		}
		tail = newNode

		lists[index] = lists[index].Next
		if lists[index] == nil {
			lists = append(lists[:index], lists[index+1:]...)
		}
	}

	if len(lists) == 1 {
		if head == nil {
			head = lists[0]
		} else {
			tail.Next = lists[0]
		}

	}
	return head
}
