package basic

import (
    "fmt"
    "testing"
)

func Test_invertTree(t *testing.T) {
    node := &TreeNode{
        Val: 4,
        Left: &TreeNode{
            Val: 2,
            Left: &TreeNode{
                Val: 1,
            },
            Right: &TreeNode{
                Val: 3,
            },
        },
        Right: &TreeNode{
            Val: 7,
            Left: &TreeNode{
                Val: 6,
            },
            Right: &TreeNode{
                Val: 9,
            },
        },
    }
    node = invertTree(node)
    fmt.Println(node)
}
