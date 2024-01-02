package basic

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	n1 := &Node{
		value: 1,
		next:  nil,
	}
	n2 := &Node{
		value: 2,
		next:  nil,
	}
	n3 := &Node{
		value: 3,
		next:  nil,
	}
	n1.next = n2
	n2.next = n3

	fmt.Println("反转之前")
	printLinkedList(n1)
	n1 = ReverseLinkedList(n1)
	fmt.Println("反转之后")
	printLinkedList(n1)
}

func printLinkedList(node *Node) {
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}
