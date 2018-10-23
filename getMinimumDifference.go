package main

func getTree(node *TreeNode, res *[]int) {

	if node.Left != nil {
		getTree(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		getTree(node.Right, res)
	}

}

func getMinimumDifference(root *TreeNode) int {
	res := []int{}
	getTree(root, &res)
	min := 1 << 32
	for i := 1; i < len(res); i++ {
		if res[i]-res[i-1] < min {
			min = res[i] - res[i-1]
		}
	}
	return min
}
