package basic

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_binaryFindLeft(t *testing.T) {
	var maxLen int64 = 20
	var maxValue int64 = 600
	times := 1000000
	succeed := true
	for i := 0; i < times; i++ {
		orderedArr := genOrderedArr(maxLen, maxValue)
		num := rand.Int63n(maxValue) - rand.Int63n(maxValue)
		if check(orderedArr, num) != binaryFindLeft(orderedArr, num) {
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

func check(arr []int64, num int64) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= num {
			return i
		}
	}
	return -1
}

func genOrderedArr(maxLen int64, maxValue int64) []int64 {
	arrLen := rand.Int63n(maxLen)
	arr := make([]int64, arrLen, arrLen)
	for i := 0; int64(i) < arrLen; i++ {
		arr[i] = rand.Int63n(maxValue)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
