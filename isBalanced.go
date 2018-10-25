package main

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if int(math.Abs(float64(getDepth(root.Left))-float64(getDepth(root.Right)))) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := getDepth(node.Left)
	right := getDepth(node.Right)
	return int(math.Max(float64(left), float64(right))) + 1
}
