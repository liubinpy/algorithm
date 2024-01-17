package basic

import (
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	//node := &TreeNode{
	//    Val:   2,
	//    Left:  &TreeNode{Val: 1},
	//    Right: &TreeNode{Val: 3},
	//}
	//
	//node := &TreeNode{
	//    Val: 5,
	//    Left: &TreeNode{
	//        Val:   1,
	//        Left:  nil,
	//        Right: nil,
	//    },
	//    Right: &TreeNode{
	//        Val: 4,
	//        Left: &TreeNode{
	//            Val:   3,
	//            Left:  nil,
	//            Right: nil,
	//        },
	//        Right: &TreeNode{
	//            Val:   6,
	//            Left:  nil,
	//            Right: nil,
	//        },
	//    },
	//}

	//node := &TreeNode{
	//	Val: 0,
	//	Left: &TreeNode{
	//		Val:   -1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: nil,
	//}

	node := &TreeNode{
		Val: 32,
		Left: &TreeNode{
			Val: 26,
			Left: &TreeNode{
				Val:  19,
				Left: nil,
				Right: &TreeNode{
					Val:   27,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  47,
			Left: nil,
			Right: &TreeNode{
				Val:   56,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isValidBST(node))
}
