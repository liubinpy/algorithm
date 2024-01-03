package basic

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	s := MyStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Size())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
