package basic

import (
	"fmt"
	"testing"
)

func TestMyQueue(t *testing.T) {
	queue := MyQueue{}
	queue.Push(1)
	queue.Push("2")
	queue.Push([]string{"haha"})
	fmt.Println(queue.Size())
	fmt.Println(queue.Peek())
	fmt.Println("====")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Size())
}
