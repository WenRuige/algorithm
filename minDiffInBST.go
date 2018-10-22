package main


func minDiffInBST(root *TreeNode) int {
	res := []int{}
	getTrees(root, &res)
	min := 1 << 32
	for i := 1; i < len(res); i++ {
		if res[i]-res[i-1] < min {
			min = res[i] - res[i-1]
		}
	}
	return min
}

func getTrees(node *TreeNode, res *[]int) {
	if node.Left != nil {
		getTrees(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		getTrees(node.Right, res)
	}
}