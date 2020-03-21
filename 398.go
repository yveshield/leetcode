package main

import (
	"crypto/rand"
	"math/big"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	var stack []int
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			stack = append(stack, i)
		}
	}
	if len(stack) == 1 {
		return stack[0]
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(len(stack))))

	return stack[result.Int64()]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
