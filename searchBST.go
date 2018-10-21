package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 如果找到了就返回这个跟
func searchBST(root *TreeNode, val int) *TreeNode {
	for {
		if root == nil {
			return nil
		}

		if val == root.Val {
			return root
		}
		if val < root.Val {
			root = root.Left
			continue
		}

		if val > root.Val {
			root = root.Right
			continue
		}
	}
}