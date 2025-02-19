package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l3head := &ListNode{}
	curr := l3head

	if list1 == nil && list2 == nil {
		return nil
	}
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curr.Val = list2.Val
			list2 = list2.Next
		} else {
			curr.Val = list1.Val
			list1 = list1.Next
		}
		curr.Next = &ListNode{}
		curr = curr.Next
	}
	if list1 == nil && list2 != nil {
		curr.Val = list2.Val
		curr.Next = list2.Next
	} else if list2 == nil && list1 != nil {
		curr.Val = list1.Val
		curr.Next = list1.Next
	}
	return l3head
}
