package basic

// 问：给定一个有序的数组，查找里数组中是否有某数字
// [1, 2, 2, 3, 6, 7, 11, 12, 13, 19, 22, 25]

func binarySearch(arr []int64, searchNum int64) bool {

	if arr == nil || len(arr) == 0 {
		return false
	}

	l := 0
	r := len(arr) - 1

	for l <= r {
		if r-l == 1 { // 临近时
			if arr[l] == searchNum || arr[r] == searchNum {
				return true
			} else {
				return false
			}
		}

		mid := (r + l) / 2

		if arr[mid] == searchNum {
			return true
		} else if arr[mid] < searchNum {
			l = mid
		} else {
			r = mid
		}
	}

	return false
}
