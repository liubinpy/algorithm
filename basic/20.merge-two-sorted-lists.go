package basic

// 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 != nil {
		return list2
	}

	if list2 == nil && list1 != nil {
		return list1
	}

	if list1 == nil && list2 == nil {
		return nil
	}

	// 先找到一个链表头小的链表，把数都合到小头上
	small := list1
	big := list2
	if list1.Val >= list2.Val {
		small = list2
		big = list1
	}
	// 记住一下head
	head := small

	var pre, nextBig *ListNode
	for small != nil && big != nil {
		if small.Val <= big.Val {
			pre = small
			small = small.Next
			if small == nil {
				pre.Next = big
			}
		} else {
			pre.Next = big
			nextBig = big.Next
			big.Next = small
			pre = big
			big = nextBig
		}
	}
	return head

}
