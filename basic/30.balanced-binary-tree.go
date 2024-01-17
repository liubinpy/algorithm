package basic

import "math"

// 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	return balance(root).isBalance
}

type Info struct {
	isBalance bool
	height    int
}

func balance(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			isBalance: true,
			height:    0,
		}
	}

	leftInfo := balance(root.Left)
	rightInfo := balance(root.Right)

	isBalance := true
	height := 0

	if !leftInfo.isBalance || !rightInfo.isBalance {
		isBalance = false
	}

	if leftInfo.isBalance && rightInfo.isBalance {
		if math.Abs(float64(leftInfo.height-rightInfo.height)) > 1 {
			isBalance = false
		}
	}

	height = max(leftInfo.height, rightInfo.height) + 1

	return &Info{
		isBalance: isBalance,
		height:    height,
	}
}
