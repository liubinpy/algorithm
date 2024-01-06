package basic

import (
	"fmt"
	"testing"
)

func TestSumLists(t *testing.T) {
	l1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  8,
		Next: nil,
	}
	//l3 := &ListNode{
	//	Val:  3,
	//	Next: nil,
	//}
	//l4 := &ListNode{
	//	Val:  9,
	//	Next: nil,
	//}
	l1.Next = l2
	//l2.Next = l3
	//l3.Next = l4

	l11 := &ListNode{
		Val:  1,
		Next: nil,
	}
	//l22 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//l33 := &ListNode{
	//	Val:  9,
	//	Next: nil,
	//}
	//l11.Next = l22
	//l22.Next = l33

	newL := addTwoNumbers(l1, l11)

	for newL != nil {
		fmt.Println(newL.Val)
		newL = newL.Next
	}

}
