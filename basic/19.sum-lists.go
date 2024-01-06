package basic

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//  先找出一个长的链表
	var short *ListNode
	var long *ListNode
	if lenList(l1) >= lenList(l2) {
		long = l1
		short = l2
	} else {
		long = l2
		short = l1
	}

	// 记录长链表的开始
	longStart := long
	longEnd := findEnd(long)
	// 进位

	var carry int
	for short != nil {
		s := short.Val + long.Val
		if carry != 0 { // 如果有进位则+1
			s += 1
			carry = 0
		}

		i := s % 10

		long.Val = i
		// 说明有进位
		if i != s {
			carry = 1
		}

		short = short.Next
		long = long.Next
	}

	for long != nil && carry == 1 {

		s := long.Val + carry
		carry = 0
		i := s % 10

		long.Val = i
		// 说明有进位
		if i != s {
			carry = 1
		}
		long = long.Next
	}

	if long == nil && carry == 1 {
		longEnd.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
		carry = 0
	}

	return longStart
}

func findEnd(l *ListNode) *ListNode {
	for l.Next != nil {
		l = l.Next
	}
	return l
}

func lenList(head *ListNode) int {
	count := 1
	for head.Next != nil {
		count++
		head = head.Next
	}
	return count
}
