package basic

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 相同的树  https://leetcode.cn/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	return p.Val == q.Val && isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)

}
