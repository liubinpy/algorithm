package basic

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {

    result := make([][]int, 0)
    if root == nil {
        return result
    }
    items := make([]int, 0)
    isPathSum(root, targetSum, &items, &result)

    return result
}

func isPathSum(root *TreeNode, targetSum int, items *[]int, result *[][]int) {
    if root.Left == nil && root.Right == nil {
        if sum(*items)+root.Val == targetSum {
            *items = append(*items, root.Val)
            res := make([]int, 0, len(*items))
            res = append(res, *items...)
            *result = append(*result, res)
            *items = (*items)[:len(*items)-1]
        }
    }

    if root.Left != nil {
        *items = append(*items, root.Val)
        isPathSum(root.Left, targetSum, items, result)
        *items = (*items)[:len(*items)-1] // 删除元素
    }

    if root.Right != nil {
        *items = append(*items, root.Val)
        isPathSum(root.Right, targetSum, items, result)
        *items = (*items)[:len(*items)-1] // 删除元素
    }
}

func sum(arr []int) int {
    s := 0
    for _, a := range arr {
        s += a
    }
    return s
}
