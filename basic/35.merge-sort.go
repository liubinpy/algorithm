package basic

// 归并排序

func mergeSort(arr []int) []int {
    if arr == nil {
        return nil
    }

    mSort(arr[0:], 0, len(arr)-1)

    return arr
}

func mSort(arr []int, begin, end int) []int {
    // 当前切片只有一个元素的时候
    if begin == end || end > len(arr)-1 {
        return arr
    }

    mid := (begin + end) / 2

    mSort(arr, begin, mid)
    mSort(arr, mid+1, end)

    merge(arr, begin, end, mid)

    return arr

}

func merge(arr []int, begin, end, mid int) {
    pt1 := begin
    pt2 := mid + 1
    tmp := make([]int, 0, end-begin)

    for pt1 <= mid && pt2 <= end {
        if arr[pt1] < arr[pt2] {
            tmp = append(tmp, arr[pt1])
            pt1++
        } else {
            tmp = append(tmp, arr[pt2])
            pt2++
        }
    }

    if pt1 <= mid {
        tmp = append(tmp, arr[pt1:mid+1]...)
    }

    if pt2 <= end {
        tmp = append(tmp, arr[pt2:end+1]...)
    }

    for i, v := range tmp {
        arr[begin+i] = v
    }
}
