package basic

// 问: 有序数组中找到>=num最左的位置，比如 [1, 2, 2, 3, 4, 6, 7],比如mun是2，应该返回1，num是3，返回3，如果是num是5，则返回5，也就是6的
// 位置

func binaryFindLeft(arr []int64, num int64) int {
	if len(arr) == 0 || arr == nil {
		return -1
	}
	L := 0
	R := len(arr) - 1
	t := -1
	for L <= R {
		mid := (L + R) / 2
		if arr[mid] >= num {
			t = mid
			R = mid - 1
		} else if arr[mid] < num {
			L = mid + 1
		}
	}
	return t
}
