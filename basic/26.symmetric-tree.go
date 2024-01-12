package basic

// 对称二叉树 https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {
	return compare(root, root)
}

func compare(left *TreeNode, right *TreeNode) bool {
	if (left != nil && right == nil) || (left == nil && right != nil) {
		return false
	}

	if left == nil && right == nil {
		return true
	}

	return left.Val == right.Val && compare(left.Right, right.Left) && compare(left.Left, right.Right)
}
