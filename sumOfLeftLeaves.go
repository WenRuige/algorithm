package main

var sum int

func getLeftLeaf(node *TreeNode) {
	if node != nil {
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				sum = sum + node.Left.Val
			}
			getLeftLeaf(node.Left)
		}
		if node.Right != nil {
			getLeftLeaf(node.Right)
		}
	}
}
func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	getLeftLeaf(root)
	return sum
}
