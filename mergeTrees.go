package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
//	if nil == t1 {
//		return t2
//	}
//	if nil == t2 {
//		return t1
//	}
//
//	return &TreeNode{Val: t1.Val + t2.Val, Left: mergeTrees(t1.Left, t2.Left), Right: mergeTrees(t1.Right, t2.Right)}
//}

func nextDepthNode(nodeSlice []*TreeNode) (nextNodeSlice []*TreeNode) {
	for _, node := range nodeSlice {
		if nil != node.Left {
			nextNodeSlice = append(nextNodeSlice, node.Left)
		}
		if nil != node.Right {
			nextNodeSlice = append(nextNodeSlice, node.Right)
		}
	}

	//fmt.Println(len(nextNodeSlice))

	return nextNodeSlice
}

func maxDepth(root *TreeNode) int {
	var maxDepth int
	if nil == root {
		return 0
	}
	// 使用一个结构来存储
	nextNodeSlice := []*TreeNode{root}
	for len(nextNodeSlice) > 0 {
		maxDepth++
		nextNodeSlice = nextDepthNode(nextNodeSlice)
	}

	return maxDepth
}

func main() {
	t1 := new(TreeNode)
	t1.Val = 1
	t1.Left = new(TreeNode)
	t1.Left.Val = 2
	t1.Right = new(TreeNode)
	t1.Right.Val = 3

	t2 := new(TreeNode)
	t2.Val = 1
	t2.Left = new(TreeNode)
	t2.Left.Val = 2
	t2.Right = new(TreeNode)
	t2.Right.Val = 3
	res := maxDepth(t1)
	fmt.Println(res)
}
