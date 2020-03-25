package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := ^int(^uint(0) >> 1)
	maxSide(root, &ans)
	return ans
}

func maxSide(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := maxSide(root.Left, ans)
	right := maxSide(root.Right, ans)
	max := left
	if right > max {
		max = right
	}
	if left+right+root.Val > *ans {
		*ans = left + right + root.Val
	}
	if max+root.Val < 0 {
		return 0
	}
	return max + root.Val
}
