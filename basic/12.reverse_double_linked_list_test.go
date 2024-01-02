package basic

import (
    "fmt"
    "testing"
)

func TestReverseDoubleLinkedList(t *testing.T) {
    n1 := &DoubleNode{
        value: 1,
        pre:   nil,
        next:  nil,
    }
    n2 := &DoubleNode{
        value: 2,
        pre:   nil,
        next:  nil,
    }
    n3 := &DoubleNode{
        value: 3,
        pre:   nil,
        next:  nil,
    }

    n1.next = n2
    n2.pre = n1
    n2.next = n3
    n3.pre = n2

    printDoubleLinkedList(n1)

    fmt.Println("反转后")
    n1 = ReverseDoubleLinkedList(n1)
    printDoubleLinkedList(n1)
}

func printDoubleLinkedList(node *DoubleNode) {
    var pre *DoubleNode
    for node != nil {
        fmt.Println(node.value)
        pre = node
        node = node.next
    }

    for pre != nil {
        fmt.Println(pre.value)
        pre = pre.pre
    }
}
