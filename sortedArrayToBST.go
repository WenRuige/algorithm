package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//func sortedArrayToBST(nums []int) *TreeNode {
//	return helper(nums, 0, len(nums)-1)
//}
//
//func helper(nums []int, i, j int) *TreeNode {
//	if len(nums) == 0 || i > j || i < 0 || j >= len(nums) {
//		return nil
//	}
//	if i == j {
//		return &TreeNode{Val: nums[i]}
//	}
//	mid := i + (j-i)/2
//	root := &TreeNode{Val: nums[mid]}
//	root.Left = helper(nums, i, mid-1)
//	root.Right = helper(nums, mid+1, j)
//	return root
//}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return getTree(nums, 0, len(nums)-1)

}

func getTree(nums []int, l, r int) *TreeNode {
	// left < right
	if l <= r {
		mid := (l + r) / 2
		treenode := new(TreeNode)
		treenode.Val = nums[mid]
		treenode.Left = getTree(nums, l, mid-1)
		treenode.Right = getTree(nums, mid+1, r)
		return treenode
	} else {
		return nil
	}
}

// 前序遍历二叉树
func preorder(node *TreeNode) {
	fmt.Println(node.Val)
	if node.Left != nil {
		preorder(node.Left)
	}
	if node.Right != nil {
		preorder(node.Right)
	}
}

// 中序遍历二叉树
func inorder(node *TreeNode) {
	if node.Left != nil {
		inorder(node.Left)
	}
	fmt.Println(node.Val)
	if node.Right != nil {
		inorder(node.Right)
	}
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)

	preorder(res)
	fmt.Println(res)
}
