package basic

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)

	// 创建一个list
	l := list.New()
	l.PushFront(root)

	for l.Len() != 0 {
		// 取出本层的长度
		r := make([]int, 0)
		n := l.Len()
		for i := 0; i < n; i++ {
			last := l.Back()
			l.Remove(last)
			e := last.Value.(*TreeNode)
			r = append(r, e.Val)

			if e.Left != nil {
				l.PushFront(e.Left)
			}
			if e.Right != nil {
				l.PushFront(e.Right)
			}
		}
		res = append(res, r)
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
