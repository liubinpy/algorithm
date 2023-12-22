package basic

// 问: 有序数组中找到<=num最右的位置，比如 [1, 2, 2, 3, 4, 6, 7],比如mun是2，应该返回2，num是3，返回3，如果是num是5，则返回4，也就是4的
// 位置

func binaryFindRight(arr []int64, num int64) int {
	t := -1
	if arr == nil || len(arr) == 0 {
		return t

	}

	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] <= num {
			t = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return t
}
