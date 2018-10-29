package main

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	getTrees(root, &res)
	return res
}

func getTrees(node *TreeNode, res *[]int) {

	if node.Left != nil {
		getTrees(node.Left, res)
	}
	if node.Right != nil {
		getTrees(node.Right, res)
	}
	*res = append(*res, node.Val)
}
