package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_binaryFindRight(t *testing.T) {
	var maxLen int64 = 20
	var maxValue int64 = 600
	times := 1000000
	succeed := true
	for i := 0; i < times; i++ {
		orderedArr := genOrderedArr(maxLen, maxValue)
		num := rand.Int63n(maxValue) - rand.Int63n(maxValue)
		if check1(orderedArr, num) != binaryFindRight(orderedArr, num) {
			succeed = false
			break
		}
	}

	if succeed {
		fmt.Println("通过")
	} else {
		fmt.Println("未通过")
	}
}

func check1(arr []int64, num int64) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= num {
			return i
		}
	}
	return -1
}
