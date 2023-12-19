package basic

import (
    "fmt"
    "math/rand"
)

// 对数器

// lenRandomValueRandom 返回一个数组，长度在[0, maxLen)，切片中的每个值在[0, maxValue-1]范围内
func lenRandomValueRandom(maxLen int64, maxValue int64) []int64 {
    sLen := int64(rand.Float64() * float64(maxLen))
    arr := make([]int64, sLen, sLen)
    for i := 0; int64(i) < sLen; i++ {
        arr[i] = int64(rand.Float64() * float64(maxValue))
    }

    return arr
}

func isSort(arr []int64) bool {
    for i, v := range arr {
        if i == len(arr)-1 {
            return true
        }
        if v > arr[i+1] {
            return false
        }
    }
    return true
}

func testMethod() bool {
    var maxLen int64 = 50
    var maxValue int64 = 1000
    testTime := 10000
    for i := 0; i < testTime; i++ {
        arr := lenRandomValueRandom(maxLen, maxValue)
        fmt.Println(arr)
        arr2 := SelectionSort(arr)
        fmt.Println(arr2)
        if !isSort(arr2) {
            return false
        }
    }
    return true
}
