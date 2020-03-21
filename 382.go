package main

import (
	"crypto/rand"
	"math/big"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	node *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	p := this.node
	q := p.Val
	for i := 1; p != nil; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
		if result.Int64() == 0 {
			q = p.Val
		}
		p = p.Next
	}
	return q
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
