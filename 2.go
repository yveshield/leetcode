package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	j := l3
	up := 0
	for l1 != nil || l2 != nil || up > 0 {
		if l1 != nil {
			up += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			up += l2.Val
			l2 = l2.Next
		}
		l3.Val = up % 10
		up /= 10
		if l1 != nil || l2 != nil || up > 0 {
			l3.Next = &ListNode{}
			l3 = l3.Next
		}
	}
	return j

}
