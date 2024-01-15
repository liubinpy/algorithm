package basic

import (
	"fmt"
	"testing"
)

func Test_buildTree(t *testing.T) {
	//tree := buildTree([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3})
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	pp(tree)
}

// 递归序
func pp(node *TreeNode) {
	if node == nil {
		return
	}

	//fmt.Println(node.Val)
	pp(node.Left)
	fmt.Println(node.Val)
	pp(node.Right)
}
