package basic

// 路径总和 https://leetcode.cn/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    var has bool
    isSum(root, 0, targetSum, &has)
    return has
}

func isSum(node *TreeNode, pre int, targetSum int, has *bool) {
    if node.Left == nil && node.Right == nil {
        if node.Val+pre == targetSum {
            *has = true
            return
        }
    }

    if node.Left != nil {
        isSum(node.Left, pre+node.Val, targetSum, has)
    }
    if *has {
        return
    }
    if node.Right != nil {
        isSum(node.Right, pre+node.Val, targetSum, has)
    }

}
