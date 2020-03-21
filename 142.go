package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	fast = head
	for fast != slow && slow.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
