package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := root.Left
	right := root.Right
	if right != nil {
		root.Left = invertTree(right)
	} else {
		root.Left = nil
	}
	if left != nil {
		root.Right = invertTree(left)
	} else {
		root.Right = nil
	}
	return root
}
