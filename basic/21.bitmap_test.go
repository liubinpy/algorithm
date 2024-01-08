package basic

import (
	"fmt"
	"testing"
)

func TestNewBitmap(t *testing.T) {
	b := NewBitmap(99999999999)
	b.Add(18211111111)
	b.Add(18211111112)
	b.Add(18211111113)
	fmt.Println(b.Contains(18211111111))
	b.Pop(18211111111)
	fmt.Println(b.Contains(18211111111))
	fmt.Println(b.Contains(18211111112))
}
