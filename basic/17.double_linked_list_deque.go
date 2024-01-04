package basic

// 双向链表实现双端队列
type dNode struct {
	value any
	pre   *dNode
	next  *dNode
}

type MyDeque struct {
	head *dNode
	tail *dNode
	size int64
}

func (q *MyDeque) Size() int64 {
	return q.size
}

func (q *MyDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *MyDeque) PushHead(value any) {
	item := &dNode{
		value: value,
		pre:   nil,
		next:  nil,
	}
	if q.head == nil {
		q.head = item
		q.tail = item
	} else {
		// 将新的元素的next指向head，将head指向元素的pre指向新的元素
		item.next = q.head
		q.head.pre = item
		// 将head移动到新的元素
		q.head = item
	}

	q.size++
}

func (q *MyDeque) PushTail(value any) {
	item := &dNode{
		value: value,
		pre:   nil,
		next:  nil,
	}
	if q.tail == nil {
		q.head = item
		q.tail = item
	} else {
		// 将新的元素的pre指向tail，将tail指向元素的next指向新的元素
		item.pre = q.tail
		q.tail.next = item

		q.tail = item
	}
	q.size++
}

func (q *MyDeque) PopHead() any {
	var item any
	if q.head == nil {
		return nil
	} else {
		item = q.head.value
		q.head = q.head.next
		if q.head != nil {
			q.head.pre = nil
		} else {
			q.tail = nil
		}

	}

	q.size--
	return item
}

func (q *MyDeque) PopTail() any {
	var item any
	if q.tail == nil {
		return nil
	} else {
		item = q.tail.value
		q.tail = q.tail.pre
		if q.tail != nil {
			q.tail.next = nil
		} else {
			q.head = nil
		}
	}

	q.size--
	return item
}

func (q *MyDeque) PeekHead() any {
	if q.head == nil {
		return nil
	} else {
		return q.head.value
	}
}

func (q *MyDeque) PeekTail() any {
	if q.tail == nil {
		return nil
	} else {
		return q.tail.value
	}
}
