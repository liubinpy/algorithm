package basic

// 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/description/
// 方法1：递归
// 方法2：中序
func isValidBST(root *TreeNode) bool {
	return isBST(root).isBST
}

type Info1 struct {
	isBST bool
	max   *Value
	min   *Value
}

type Value struct {
	v int
}

func isBST(root *TreeNode) *Info1 {
	if root == nil {
		return &Info1{
			isBST: true,
			max:   nil,
			min:   nil,
		}
	}

	left := isBST(root.Left)
	right := isBST(root.Right)

	isbst := true
	var maxValue, minValue *Value

	if left.max == nil && right.min == nil {
		maxValue = &Value{v: root.Val}
		minValue = &Value{v: root.Val}
	}

	if left.max != nil && right.min != nil {
		if root.Val <= left.max.v {
			isbst = false
		}
		if root.Val >= right.min.v {
			isbst = false
		}
		if !left.isBST || !right.isBST {
			isbst = false
		}

		maxValue = &Value{v: max(left.max.v, right.max.v, root.Val)}
		minValue = &Value{v: min(left.min.v, right.min.v, root.Val)}
	}

	if left.max == nil && right.min != nil {
		if root.Val >= right.min.v {
			isbst = false
		}
		if !right.isBST {
			isbst = false
		}

		maxValue = &Value{v: max(right.max.v, root.Val)}
		minValue = &Value{v: min(right.min.v, root.Val)}
	}

	if left.max != nil && right.min == nil {
		if root.Val <= left.max.v {
			isbst = false
		}
		if !left.isBST {
			isbst = false
		}

		maxValue = &Value{v: max(left.max.v, root.Val)}
		minValue = &Value{v: min(left.min.v, root.Val)}
	}

	return &Info1{
		isBST: isbst,
		max:   maxValue,
		min:   minValue,
	}

}
