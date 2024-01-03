package basic

// 单向链表实现队列，先进先出

type QNode struct {
	value any
	next  *QNode
}

type MyQueue struct {
	head *QNode
	tail *QNode
	size int64
}

func (q *MyQueue) Size() int64 {
	return q.size
}

func (q *MyQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *MyQueue) Push(value any) {
	item := &QNode{
		value: value,
		next:  nil,
	}
	if q.head == nil {
		q.head = item
		q.tail = item
	} else {
		q.tail.next = item
		q.tail = item
	}
	q.size++
}

func (q *MyQueue) Pop() any {
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

func (q *MyQueue) Peek() any {
	if q.head == nil {
		return nil
	} else {
		return q.head.value
	}

}
