package basic

// 反转双向链表

type DoubleNode struct {
	value int64
	pre   *DoubleNode
	next  *DoubleNode
}

func ReverseDoubleLinkedList(head *DoubleNode) *DoubleNode {
	var next *DoubleNode
	var pre *DoubleNode
	for head != nil {
		next = head.next
		head.next = pre
		head.pre = next
		pre = head
		head = next
	}
	return pre
}
