package basic

import "fmt"

// InsertionSort 插入排序
func InsertionSort(s []int64) []int64 {
	for i := 0; i < len(s); i++ {
		for ; i > 0 && s[i-1] > s[i]; i-- {
			s[i-1], s[i] = s[i], s[i-1]
			fmt.Println(s)
		}
	}
	return s
}
