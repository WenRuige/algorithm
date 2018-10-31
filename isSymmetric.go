package main

// 是否是对称二叉树

/*

如果左右子树有一个节点为空,那么不是对称二叉树


*/

func isSymmetric(root *TreeNode) bool {
	return symmetry(root, root)
}

// 对称
func symmetry(left *TreeNode, right *TreeNode) bool {
	//
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return symmetry(left.Left, right.Right) && symmetry(left.Right, right.Left)

}
