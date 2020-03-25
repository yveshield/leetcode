package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var pre *TreeNode
	for root != nil {
		if root.Left != nil {
			pre = root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root
			tmp := root
			root = root.Left
			tmp.Left = nil
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res

}
