package main


// 二叉搜索树的性质

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	// 如果根节点为空 return
	if root == nil {
		return nil
	}
	// 如果根节点的值小于l,说明 它的左子树都不符合条件,因为左子树要小于根节点
	if root.Val < L {
		return trimBST(root.Right, L, R)
		//如果根节点的值大于R,说明 它的右子树都不符合条件,因为右子树要大于根节点
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	} else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}

}
