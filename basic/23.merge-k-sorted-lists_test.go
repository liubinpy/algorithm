package basic

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	l1 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  8,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l51 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l52 := &ListNode{
		Val:  22,
		Next: nil,
	}
	l53 := &ListNode{
		Val:  99,
		Next: nil,
	}
	l5.Next = l51
	l51.Next = l52
	l52.Next = l53

	//list := []*ListNode{nil, nil}
	list := []*ListNode{l2, l5, l1, l3, l4}
	one := mergeKLists(list)
	for one != nil {
		fmt.Println("=", one.Val)
		one = one.Next
	}
}
