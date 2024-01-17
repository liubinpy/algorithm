package basic

import (
	"fmt"
	"testing"
)

func Test_balance(t *testing.T) {
	//node := &TreeNode{
	//    Val: 1,
	//    Left: &TreeNode{
	//        Val: 2,
	//        Left: &TreeNode{
	//            Val: 3,
	//            Left: &TreeNode{
	//                Val:   4,
	//                Left:  nil,
	//                Right: nil,
	//            },
	//            Right: &TreeNode{
	//                Val:   4,
	//                Left:  nil,
	//                Right: nil,
	//            },
	//        },
	//        Right: &TreeNode{
	//            Val:   3,
	//            Left:  nil,
	//            Right: nil,
	//        },
	//    },
	//    Right: &TreeNode{
	//        Val:   2,
	//        Left:  nil,
	//        Right: nil,
	//    },
	//}

	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isBalanced(node))
}
