package basic

// 选择排序，每次都找到一个最小的
//[2 5 9 10 1 5 3]
//[1 5 9 10 2 5 3]
//[1 2 9 10 5 5 3]
//[1 2 3 10 5 5 9]
//[1 2 3 5 10 5 9]
//[1 2 3 5 5 10 9]
//[1 2 3 5 5 9 10]
//[1 2 3 5 5 9 10]

func SelectionSort(s []int64) []int64 {
	for i := 0; i < len(s); i++ {
		curMinIdx := i
		for j := i; j < len(s); j++ {
			if s[j] < s[curMinIdx] {
				curMinIdx = j
			}
		}
		s[i], s[curMinIdx] = s[curMinIdx], s[i]
	}
	return s
}
