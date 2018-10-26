package main

func getTrees(node *TreeNode, res *[]int) {
	if node.Left != nil {
		getTrees(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		getTrees(node.Right, res)
	}
}

func kthSmallestTree(root *TreeNode, k int) int {
	res := []int{}
	getTrees(root, &res)
	return res[k-1]
}
