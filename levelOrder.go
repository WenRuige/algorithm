package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := []int{}
	final := [][]int{}

	parentSize := 1
	var nodes = []TreeNode{*root}
	level := 0
	for len(nodes) != 0 {
		// 根节点
		node := nodes[0]
		nodes = nodes[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			level++
			nodes = append(nodes, *node.Left)
		}

		if node.Right != nil {
			level++
			nodes = append(nodes, *node.Right)
		}
		parentSize--
		if parentSize == 0 {
			parentSize = level
			level = 0
			final = append(final, result)
			result = []int{}
		}

	}

	fmt.Println(final)

	return final
}

func main() {

	node := new(TreeNode)
	node.Val = 1
	node.Left = new(TreeNode)
	node.Right = new(TreeNode)
	node.Left.Val = 2
	node.Right.Val = 3

	node.Left.Left = new(TreeNode)
	node.Left.Left.Val = 4

	res := levelOrder(node)
	fmt.Println(res)
}
