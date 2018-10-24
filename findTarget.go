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
func findTarget(root *TreeNode, k int) bool {
	res := []int{}
	getTree(root, &res)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i]+res[j] == k {
				return true
			}
		}
	}
	return false
}
