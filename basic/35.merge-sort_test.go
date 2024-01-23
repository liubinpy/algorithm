package basic

import (
    "fmt"
    "testing"
)

func Test_mergeSort(t *testing.T) {
    arr := []int{4, 3, 8, 1, 6, 5, 2, 9, 0, 100, 7}
    mergeSort(arr)
    fmt.Println(arr)
}
