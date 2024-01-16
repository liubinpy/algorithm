package basic

import (
	"fmt"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	node := &TreeNode{
		Val: 3,
	}
	node.Left = &TreeNode{
		Val: 9,
	}
	node.Right = &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 15,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	orderBottom := levelOrderBottom(node)

	fmt.Printf("%#v", orderBottom)
}
