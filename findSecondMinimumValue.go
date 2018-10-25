package main

import "sort"

func findSecondMinimumValue(root *TreeNode) int {

	res := []int{}
	getTrees(root, &res)
	sort.Ints(res)
	min := res[0]
	for i := 0; i < len(res); i++ {
		if min < res[i] {
			return res[i]
		}
	}
	return -1

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
