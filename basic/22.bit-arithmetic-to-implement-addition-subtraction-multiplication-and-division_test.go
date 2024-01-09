package basic

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(Add(44, 99))
}

func TestDiv(t *testing.T) {
	fmt.Println(div(-2147483648, -3))
}

func TestMulti(t *testing.T) {
	fmt.Println(Multi(18, 19))
}

func TestSub(t *testing.T) {
	fmt.Println(Sub(10, 1))
}

func Test_divide(t *testing.T) {
	fmt.Println(divide(7, -3))
}
