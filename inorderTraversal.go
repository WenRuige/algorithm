package main

// 二叉树的中序遍历
func getTrees(node *TreeNode, res *[]int) {
	if node.Left != nil {
		getTrees(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		getTrees(node.Right, res)
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	getTrees(root, &res)
	return res
}
