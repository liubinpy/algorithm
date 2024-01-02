package basic

// 问: 给定一个无序数组，相邻的数不相等，那么就会产生很多个局部最小的值，找到任意一个局部最小的值的位置返回
// 局部最小：一个数满足小于相邻的左边的数和小于相邻的右边的数，如果0位置上的数小于1位置上的数，那么就返回0，如果n-1位置上面的数小于n-2位置上面
// 的数就返回n-1  例如：arr=[5,2,1,6,4,1,10,9,20]， 返回：2或1或9任意一个都可以
// 3 1 2 1 3
func localMinimum(arr []int64) int {
	aLen := len(arr)
	if arr == nil || aLen == 0 {
		return -1
	}
	if aLen == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[aLen-1] < arr[aLen-2] {
		return aLen - 1
	}

	l := 0
	r := len(arr) - 1
	mid := -1

	for l <= r {
		mid = (l + r) / 2
		if arr[mid] < arr[mid+1] && arr[mid] < arr[mid-1] {
			return mid
		}

		if arr[mid] < arr[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return mid
}
