package basic

// BubbleSort 冒泡排序
func BubbleSort(s []int64) []int64 {
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
