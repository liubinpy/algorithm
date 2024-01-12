package basic

import "fmt"

type BNode struct {
	value int
	left  *BNode
	right *BNode
}

// 递归序
func ff(node *BNode) {
	if node == nil {
		return
	}

	ff(node.left)
	ff(node.right)
}

// 先序，头左右
func pre(node *BNode) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	pre(node.left)
	pre(node.right)
}

// 中序，左头右
func middle(node *BNode) {
	if node == nil {
		return
	}
	middle(node.left)
	fmt.Println(node.value)
	middle(node.right)
}

// 后序：左右头
func pos(node *BNode) {
	if node == nil {
		return
	}
	pos(node.left)
	pos(node.right)
	fmt.Println(node.value)
}
