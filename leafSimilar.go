package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1 := []int{}
	res2 := []int{}
	getTree(root1, &res1)
	getTree(root2, &res2)

	if len(res1) != len(res2) {
		return false
	}

	for i := 0; i < len(res1); i++ {
		if res1[i] != res2[i] {
			return false
		}
	}
	return true

}

func getTree(node *TreeNode, res *[]int) {
	if node.Left == nil && node.Right == nil {
		*res = append(*res, node.Val)
	} else {
		if node.Left != nil {
			getTree(node.Left, res)
		}
		if node.Right != nil {
			getTree(node.Right, res)
		}
	}

}
