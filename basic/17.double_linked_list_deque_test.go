package basic

import (
	"fmt"
	"testing"
)

func TestMyDeque(t *testing.T) {
	m := MyDeque{}

	m.PushHead(1)
	m.PushHead(2)
	m.PushTail("aaa")
	m.PushTail("bbb")

	fmt.Println(m.size)
	fmt.Println(m.PeekHead())
	fmt.Println(m.size)

	fmt.Println("====")
	fmt.Println(m.PopTail())
	fmt.Println(m.PopTail())
	fmt.Println(m.PopTail())
	fmt.Println(m.PopTail())
	fmt.Println(m.size)
	fmt.Println(m.PopTail())
	fmt.Println(m.PopHead())
	fmt.Println("============")
	m.PushHead(1)
	m.PushHead(2)
	m.PushHead(3)
	fmt.Println(m.PopHead())
	fmt.Println(m.PopHead())
	fmt.Println(m.PopHead())
	fmt.Println(m.PopHead())
	fmt.Println(m.PopHead())
}
