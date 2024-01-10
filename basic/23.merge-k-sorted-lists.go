package basic

import (
	"container/heap"
)

// 合并K个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/description/

type MyHeap []*ListNode

func (m MyHeap) Len() int {
	return len(m)
}

func (m MyHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m MyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MyHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *MyHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := &MyHeap{}
	for _, list := range lists {
		if list != nil {
			m.Push(list)
		}
	}
	heap.Init(m)
	var pre *ListNode
	var tail *ListNode
	for m.Len() > 0 {
		// 取出最上面的链表的第一个
		item := heap.Pop(m).(*ListNode)
		if pre == nil {
			pre = item
			tail = item
		} else {
			tail.Next = item
			tail = item
		}
		if item.Next != nil {
			heap.Push(m, item.Next)
		}
	}
	return pre
}
