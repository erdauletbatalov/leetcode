package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodeArr := []int{}
	curr := head
	for {
		nodeArr = append(nodeArr, curr.Val)
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}
	newHead := &ListNode{}
	curr = newHead
	for i := len(nodeArr) - 1; i >= 0; i-- {
		curr.Val = nodeArr[i]
		if i == 0 {
			break
		}
		curr.Next = &ListNode{}
		curr = curr.Next
	}
	return newHead
}
