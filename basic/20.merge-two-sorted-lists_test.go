package basic

import (
    "fmt"
    "testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
    l1 := &ListNode{
        Val:  1,
        Next: nil,
    }
    l2 := &ListNode{
        Val:  1,
        Next: nil,
    }
    l3 := &ListNode{
        Val:  4,
        Next: nil,
    }
    l4 := &ListNode{
        Val:  5,
        Next: nil,
    }
    l5 := &ListNode{
        Val:  7,
        Next: nil,
    }
    l1.Next = l2
    l2.Next = l3
    l3.Next = l4
    l4.Next = l5

    l11 := &ListNode{
        Val:  2,
        Next: nil,
    }
    l22 := &ListNode{
        Val:  3,
        Next: nil,
    }
    l33 := &ListNode{
        Val:  6,
        Next: nil,
    }

    l11.Next = l22
    l22.Next = l33

    lists := mergeTwoLists(l1, l11)
    for lists != nil {
        fmt.Println(lists.Val)
        lists = lists.Next
    }
}
