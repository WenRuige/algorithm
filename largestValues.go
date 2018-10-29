package main

import "math"

func largestValues(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	// 声明一个存储空间
	queue := make([]*TreeNode, 1, 1024)
	// 0为根节点
	queue[0] = root
	// n = 1 二叉树层数
	n := 1
	max := math.MinInt32
	for len(queue) > 0 {
		n--
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if max < node.Val {
			max = node.Val
		}
		// 层数
		if n == 0 {
			n = len(queue)
			result = append(result, max)
			max = math.MinInt32
		}
	}
	return result

}
