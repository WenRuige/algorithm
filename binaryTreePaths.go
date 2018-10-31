package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	getTreePath(root, strconv.Itoa(root.Val), &res)
	return res
}

func getTreePath(node *TreeNode, path string, res *[]string) {
	if node.Left == nil && node.Right == nil {
		*res = append(*res, path)
	}
	//fmt.Println(node.Val)
	if node.Left != nil {
		getTreePath(node.Left, path+"->"+strconv.Itoa(node.Left.Val), res)
	}
	if node.Right != nil {
		getTreePath(node.Right, path+"->"+strconv.Itoa(node.Right.Val), res)
	}

}
