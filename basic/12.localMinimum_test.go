package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_localMinimum(t *testing.T) {
	var maxLen int64 = 100
	var maxValue int64 = 200
	times := 1000000
	for i := 0; i < times; i++ {
		arr := genUnOrderedArr(maxLen, maxValue)
		index := localMinimum(arr)
		if !checkMini(arr, index) {
			fmt.Println(arr)
			fmt.Println("失败咯")
			break
		}
	}

}

func checkMini(arr []int64, index int) bool {
	if len(arr) == 1 {
		return index == 0
	}
	if index == 0 {
		return arr[0] < arr[1]
	}
	if index == len(arr)-1 {
		return arr[len(arr)-1] < arr[len(arr)-2]
	}

	return arr[index] < arr[index+1] && arr[index] < arr[index-1]
}

func genUnOrderedArr(maxLen int64, maxValue int64) []int64 {
	arrLen := rand.Int63n(maxLen) + 1
	arr := make([]int64, arrLen, arrLen)
	arr[0] = rand.Int63n(maxValue)
	for i := 1; int64(i) < arrLen; i++ {
		arr[i] = rand.Int63n(maxValue)
		for arr[i] == arr[i-1] {
			arr[i] = rand.Int63n(maxValue)
		}
	}
	return arr
}
