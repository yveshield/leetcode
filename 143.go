package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	n := 0
	tail := new(ListNode)
	var p *ListNode
	p = head
	for p != nil {
		tail.Val = p.Val
		p = p.Next
		if p == nil {
			break
		}
		tmp := new(ListNode)
		tmp.Next = tail
		tail = tmp
		n++
	}
	if n < 2 {
		return
	}
	direction := 0
	relist := new(ListNode)
	var q *ListNode
	q = relist
	p = head
	for i := 0; i <= n; i++ {
		if direction == 0 {
			q.Val = p.Val
			p = p.Next
		} else {
			q.Val = tail.Val
			tail = tail.Next
		}
		if i == n {
			break
		}
		q.Next = new(ListNode)
		q = q.Next
		direction = 1 - direction
	}
	head.Next = relist.Next
}
