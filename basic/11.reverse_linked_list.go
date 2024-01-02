package basic

// 反转单向链表

type Node struct {
	value int64
	next  *Node
}

func ReverseLinkedList(head *Node) *Node {
	var next *Node
	var pre *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}
