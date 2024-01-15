package basic

// 105. 从前序与中序遍历序列构造二叉树
// 思路就是，前序序列是用来找到头部的位置，中序序列是用来找到左右的元素
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}

	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}

	//  前序的第一个元素肯定是二叉树的头
	head := &TreeNode{
		Val: preorder[0],
	}

	// 找到head在中序的索引
	var index int
	for i, v := range inorder {
		if v == head.Val {
			index = i
			break
		}
	}

	head.Left = buildTree(preorder[1:index+1], inorder[0:index])
	head.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return head
}
