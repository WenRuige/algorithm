package main

// 二叉树的前序遍历

func getTrees(node *TreeNode, res *[]int) {
	*res = append(*res, node.Val)
	if node.Left != nil {
		getTrees(node.Left, res)
	}
	if node.Right != nil {
		getTrees(node.Right, res)
	}
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	getTrees(root, &res)
	return res
}
