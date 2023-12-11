package basic

// 有一个数组arr,求其n到m的和，m>=n

// n2mSumWay1 解法1，构建一个二维数组
func n2mSumWay1(arr []int64, n, m int64) int64 {
	table := make([][]int64, 0, len(arr))
	for i1, v1 := range arr {
		row := make([]int64, 0, len(arr))
		for i2, v2 := range arr {
			if i2 == i1 {
				row = append(row, v1)
			} else if i2 > i1 {
				row = append(row, v2+row[i2-1])
			} else {
				row = append(row, 0)
			}
		}
		table = append(table, row)
	}
	if m < n {
		panic("请输入合法的m和n，m必须大于等于n")
	}

	return table[n][m]
}

// n2mSumWay2 前缀和
func n2mSumWay2(arr []int64, n, m int64) int64 {
	preArr := make([]int64, 0, len(arr))
	for i, v := range arr {
		if i == 0 {
			preArr = append(preArr, v)
		} else {
			preArr = append(preArr, preArr[i-1]+v)
		}
	}

	if n > m {
		panic("请输入合法的m和n，m必须大于等于n")
	}

	if n == 0 {
		return preArr[m]
	}

	return preArr[m] - preArr[n-1]
}
