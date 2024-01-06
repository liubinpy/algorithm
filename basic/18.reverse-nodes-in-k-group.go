package basic

// K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if lens(head) < k {
		return head
	}

	groupEnd := findGroupEnd(head, k)
	lastHead := head
	firstHead := groupEnd
	nextHead := groupEnd.Next
	reverse(head, groupEnd)
	if nextHead == nil {
		return firstHead
	}

	groupEnd = findGroupEnd(nextHead, k)
	for groupEnd != nil {
		head = nextHead
		nextHead = groupEnd.Next
		reverse(head, groupEnd)

		lastHead.Next = groupEnd
		lastHead = head
		if nextHead == nil {
			return firstHead
		}
		groupEnd = findGroupEnd(nextHead, k)
	}
	return firstHead

}

func lens(head *ListNode) int {
	count := 1
	for head.Next != nil {
		count++
		head = head.Next
	}
	return count
}

func findGroupEnd(start *ListNode, k int) *ListNode {
	end := start
	for k > 1 {
		end = end.Next
		k--
		if end == nil {
			return end
		}
	}
	return end
}

func reverse(start, end *ListNode) *ListNode {
	end = end.Next
	beginStart := start
	var pre, next *ListNode

	for start != end {
		next = start.Next
		start.Next = pre
		pre = start
		start = next
	}
	beginStart.Next = start
	return pre
}
