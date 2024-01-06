package basic

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	//l5 := &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	//l6 := &ListNode{
	//	Val:  6,
	//	Next: nil,
	//}
	//l7 := &ListNode{
	//	Val:  7,
	//	Next: nil,
	//}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	//l4.Next = l5
	//l5.Next = l6
	//l6.Next = l7

	l1 = reverseKGroup(l1, 2)

	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}

}
