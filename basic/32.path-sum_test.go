package basic

import (
    "fmt"
    "testing"
)

func Test_hasPathSum(t *testing.T) {
    //node := &TreeNode{
    //    Val: 5,
    //    Left: &TreeNode{
    //        Val: 4,
    //        Left: &TreeNode{
    //            Val: 11,
    //            Left: &TreeNode{
    //                Val: 7,
    //            },
    //            Right: &TreeNode{
    //                Val: 2,
    //            },
    //        },
    //    },
    //    Right: &TreeNode{
    //        Val: 8,
    //        Left: &TreeNode{
    //            Val: 13,
    //        },
    //        Right: &TreeNode{
    //            Val: 4,
    //            Right: &TreeNode{
    //                Val: 1,
    //            },
    //        },
    //    },
    //}

    node := &TreeNode{
        Val: 1,
        Right: &TreeNode{
            Val: 2,
            Right: &TreeNode{
                Val: 3,
                Right: &TreeNode{
                    Val: 4,
                    Right: &TreeNode{
                        Val: 5,
                    },
                },
            },
        },
    }

    fmt.Println(hasPathSum(node, 3))
}
