package basic

// 单链表实现栈，先进后出

type SNode struct {
	value any
	next  *SNode
}

type MyStack struct {
	head *SNode
	size int64
}

func (q *MyStack) Size() int64 {
	return q.size
}

func (q *MyStack) IsEmpty() bool {
	return q.size == 0
}

func (q *MyStack) Push(value any) {
	item := &SNode{
		value: value,
		next:  nil,
	}
	if q.head == nil {
		q.head = item
	} else {
		item.next = q.head
		q.head = item
	}
	q.size++
}

func (q *MyStack) Pop() any {
	var item any
	if q.head == nil {
		return nil
	} else {
		item = q.head.value
		q.head = q.head.next
	}
	q.size--
	return item
}

func (q *MyStack) Peek() any {
	if q.head == nil {
		return nil
	} else {
		return q.head.value
	}
}
