package basic

import (
	"fmt"
	"testing"
)

func TestBNode(t *testing.T) {
	node := &BNode{
		value: 1,
	}

	node.left = &BNode{
		value: 2,
	}
	node.left.left = &BNode{
		value: 4,
	}
	node.left.right = &BNode{
		value: 5,
	}

	node.right = &BNode{
		value: 3,
	}
	node.right.left = &BNode{
		value: 6,
	}
	node.right.right = &BNode{
		value: 7,
	}

	// 先序：头左右
	pre(node)
	fmt.Println("========")
	// 中序
	middle(node)
	// 后序
	fmt.Println("========")
	pos(node)
}
