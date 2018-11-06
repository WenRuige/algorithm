package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	res := helper(nums, 0, len(nums)-1)
	return res
}

// helper 分治法解决
func helper(nums []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}
	max := left

	for i := left + 1; i <= right; i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}

	tree := new(TreeNode)
	tree.Val = nums[max]
	tree.Left = helper(nums, left, max-1)
	tree.Right = helper(nums, max+1, right)
	return tree
}

func main() {

}
